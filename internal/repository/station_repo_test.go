package repository

import (
	"context"
	"main/internal/dao"
	"main/internal/database"
	"testing"
	"time"

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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo()

	station := dao.Station{
		SiteID:   "asdadqerqw",
		Brand:    "TestBrand",
		Address:  "123 Test St",
		Postcode: "TE5 7ST",
		Location: dao.GeometryPoint{
			X: -0.0706,
			Y: 51.8018,
		},
		E5:        199.9,
		E10:       199.9,
		B7:        199.9,
		SDV:       199.9,
		CreatedAt: time.Now(),
	}

	err = repo.Insert(ctx, &station)
	assert.NoError(t, err)
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo()
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewStationRepo()
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

func writeStations(t *testing.T, repo StationRepository) {
	stations := []dao.Station{
		{
			SiteID:    "station1",
			Brand:     "BrandA",
			Address:   "Address 1",
			Postcode:  "PC1 1AA",
			Location:  dao.GeometryPoint{X: -0.1, Y: 51.5},
			E5:        150.0,
			CreatedAt: time.Now(),
		},
		{
			SiteID:    "station2",
			Brand:     "BrandB",
			Address:   "Address 2",
			Postcode:  "PC2 2BB",
			Location:  dao.GeometryPoint{X: -0.2, Y: 51.6},
			E5:        145.0,
			E10:       140.0,
			CreatedAt: time.Now(),
		},
		{
			SiteID:    "station3",
			Brand:     "BrandA",
			Address:   "Address 3",
			Postcode:  "PC3 3CC",
			Location:  dao.GeometryPoint{X: -0.3, Y: 51.7},
			B7:        130.0,
			CreatedAt: time.Now(),
		},
	}

	for _, station := range stations {
		err := repo.Insert(context.Background(), &station)
		assert.NoError(t, err)
	}
}
