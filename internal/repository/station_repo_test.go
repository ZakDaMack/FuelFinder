package repository

import (
	"context"
	"database/sql"
	"main/internal/dao"
	"main/internal/database"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo(db)

	station := dao.Station{
		SiteID:   "asdadqerqw",
		Brand:    "TestBrand",
		Address:  "123 Test St",
		Postcode: "TE5 7ST",
		Location: dao.GeometryPoint{
			X: -0.0706,
			Y: 51.8018,
		},
		E5:        sql.NullFloat64{Float64: 199.9, Valid: true},
		E10:       sql.NullFloat64{Float64: 199.9, Valid: true},
		B7:        sql.NullFloat64{Float64: 199.9, Valid: true},
		SDV:       sql.NullFloat64{Float64: 199.9, Valid: true},
		CreatedAt: time.Now(),
	}

	err = repo.Insert(ctx, &station)
	assert.NoError(t, err)
}

func TestDuplicateInsert(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo(db)

	station := dao.Station{
		ID:       uuid.New(),
		SiteID:   "asdadqerqw",
		Brand:    "TestBrand",
		Address:  "123 Test St",
		Postcode: "TEE5 7ST",
		Location: dao.GeometryPoint{
			X: -0.0706,
			Y: 51.8018,
		},
		E5:        sql.NullFloat64{Float64: 199.9, Valid: true},
		E10:       sql.NullFloat64{Float64: 199.9, Valid: true},
		B7:        sql.NullFloat64{Float64: 199.9, Valid: true},
		SDV:       sql.NullFloat64{Float64: 199.9, Valid: true},
		CreatedAt: time.Now(),
	}

	// Insert station first time
	err = repo.Insert(ctx, &station)
	assert.NoError(t, err)

	// Attempt to insert duplicate station with same site_id and created_at
	station.ID = uuid.New()
	err = repo.Insert(ctx, &station)
	assert.NoError(t, err) // Should not error due to ON CONFLICT DO NOTHING

	// Verify only one record exists in the database
	var count int64
	err = db.Model(&dao.Station{}).Where("site_id = ? AND created_at = ?", station.SiteID, station.CreatedAt).Count(&count).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestGetBrands(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo(db)
	writeStations(t, repo)

	brands, err := repo.GetBrands(ctx)
	assert.NoError(t, err)
	assert.Len(t, brands, 2)
	assert.ElementsMatch(t, []string{"BrandA", "BrandB"}, brands)
}

func TestGetStations(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo(db)
	writeStations(t, repo)

	topLeft := dao.GeometryPoint{X: -0.5, Y: 51.8}
	bottomRight := dao.GeometryPoint{X: 0.5, Y: 50.0}

	// get all stations in bounding box
	stations, err := repo.GetStations(ctx, topLeft, bottomRight, []string{}, []string{})
	assert.NoError(t, err)
	for _, station := range stations {
		assert.Contains(t, []string{"BrandA", "BrandB"}, station.Brand)
	}
}

func TestGetStationsWithFilters(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	assert.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo(db)
	writeStations(t, repo)

	topLeft := dao.GeometryPoint{X: -0.5, Y: 51.8}
	bottomRight := dao.GeometryPoint{X: 0.5, Y: 50.0}

	// get all stations in bounding box with fueltype filter
	stations, err := repo.GetStations(ctx, topLeft, bottomRight, []string{}, []string{"B7"})
	assert.NoError(t, err)
	assert.Len(t, stations, 1)
	assert.Equal(t, "station3", stations[0].SiteID)
	assert.NotNil(t, stations[0].B7)

	// get all stations in bounding box with brand filter
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandA"}, []string{})
	assert.NoError(t, err)
	assert.Len(t, stations, 2)
	for _, station := range stations {
		assert.Equal(t, "BrandA", station.Brand)
	}

	// get all stations in bounding box with brand and fueltype filter
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandB"}, []string{"E5"})
	assert.NoError(t, err)
	assert.Len(t, stations, 1)
	assert.Equal(t, "BrandB", stations[0].Brand)
	assert.Equal(t, "station2", stations[0].SiteID)
	assert.NotNil(t, stations[0].E5)

	// get only latest entry for station1
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandA"}, []string{})
	assert.NoError(t, err)
	for _, station := range stations {
		if station.SiteID == "station1" {
			assert.Equal(t, "Address 1 Updated", station.Address)
		}
	}
}

func writeStations(t *testing.T, repo StationRepository) {
	stations := []dao.Station{
		{
			SiteID:    "station1",
			Brand:     "BrandA",
			Address:   "Address 1",
			Postcode:  "PC1 1AA",
			Location:  dao.GeometryPoint{X: -0.1, Y: 51.5},
			E5:        sql.NullFloat64{Float64: 150.0, Valid: true},
			CreatedAt: time.Now(),
		},
		{
			SiteID:    "station2",
			Brand:     "BrandB",
			Address:   "Address 2",
			Postcode:  "PC2 2BB",
			Location:  dao.GeometryPoint{X: -0.2, Y: 51.6},
			E5:        sql.NullFloat64{Float64: 145.0, Valid: true},
			E10:       sql.NullFloat64{Float64: 140.0, Valid: true},
			CreatedAt: time.Now(),
		},
		{
			SiteID:    "station3",
			Brand:     "BrandA",
			Address:   "Address 3",
			Postcode:  "PC3 3CC",
			Location:  dao.GeometryPoint{X: -0.3, Y: 51.7},
			B7:        sql.NullFloat64{Float64: 130.0, Valid: true},
			CreatedAt: time.Now(),
		},
		// test same station with different timestamp
		{
			SiteID:    "station1",
			Brand:     "BrandA",
			Address:   "Address 1 Updated",
			Postcode:  "PC1 1AA",
			Location:  dao.GeometryPoint{X: -0.1, Y: 51.5},
			E5:        sql.NullFloat64{Float64: 149.0, Valid: true},
			CreatedAt: time.Now().Add(10 * time.Minute),
		},
	}

	for _, station := range stations {
		err := repo.Insert(context.Background(), &station)
		assert.NoError(t, err)
	}
}
