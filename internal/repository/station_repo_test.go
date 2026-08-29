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
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

var priceBase = time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)

func newTestRepos(t *testing.T) (StationRepository, StationPriceRepository, *gorm.DB) {
	t.Helper()
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgis/postgis:12-3.0",
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(30*time.Second)),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx)
	require.NoError(t, err)

	db, err := database.MakePostgresDB(dsn)
	require.NoError(t, err)

	return NewStationRepo(db), NewStationPriceRepo(db), db
}

func TestUpsert(t *testing.T) {
	ctx := context.Background()
	repo, _, _ := newTestRepos(t)

	station := dao.Station{
		NodeID:   "asdadqerqw",
		Brand:    "TestBrand",
		Address:  "123 Test St",
		Postcode: "TE5 7ST",
		Location: dao.GeometryPoint{
			X: -0.0706,
			Y: 51.8018,
		},
	}

	err := repo.Upsert(ctx, &station)
	assert.NoError(t, err)
}

func TestUpsertUpdatesExisting(t *testing.T) {
	ctx := context.Background()
	repo, _, db := newTestRepos(t)

	station := dao.Station{
		NodeID:   "asdadqerqw",
		Name:     "Test Station",
		Brand:    "TestBrand",
		Address:  "123 Test St",
		Postcode: "TE5 7ST",
		Location: dao.GeometryPoint{
			X: -0.0706,
			Y: 51.8018,
		},
	}

	// Insert station first time
	err := repo.Upsert(ctx, &station)
	assert.NoError(t, err)

	// Upsert the same node again with changed details
	updated := station
	updated.ID = uuid.Nil
	updated.Address = "456 New Rd"
	err = repo.Upsert(ctx, &updated)
	assert.NoError(t, err)

	// Verify only one record exists and it carries the new details
	var results []dao.Station
	err = db.Model(&dao.Station{}).Where("node_id = ?", station.NodeID).Find(&results).Error
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "456 New Rd", results[0].Address)
}

func TestGetBrands(t *testing.T) {
	ctx := context.Background()
	repo, _, _ := newTestRepos(t)
	writeStations(t, repo)

	brands, err := repo.GetBrands(ctx)
	assert.NoError(t, err)
	assert.Len(t, brands, 2)
	assert.ElementsMatch(t, []string{"BrandA", "BrandB"}, brands)
}

func TestGetStations(t *testing.T) {
	ctx := context.Background()
	repo, _, _ := newTestRepos(t)
	writeStations(t, repo)

	topLeft := dao.GeometryPoint{X: -0.5, Y: 51.8}
	bottomRight := dao.GeometryPoint{X: 0.5, Y: 50.0}

	// get all stations in bounding box
	stations, err := repo.GetStations(ctx, topLeft, bottomRight, []string{}, []string{})
	assert.NoError(t, err)
	assert.Len(t, stations, 3)
	for _, station := range stations {
		assert.Contains(t, []string{"BrandA", "BrandB"}, station.Brand)
	}
}

func TestGetStationsWithFilters(t *testing.T) {
	ctx := context.Background()
	repo, priceRepo, _ := newTestRepos(t)
	writeStations(t, repo)
	writePrices(t, priceRepo)

	topLeft := dao.GeometryPoint{X: -0.5, Y: 51.8}
	bottomRight := dao.GeometryPoint{X: 0.5, Y: 50.0}

	// get all stations in bounding box with fueltype filter (legacy name b7 maps to b7_standard)
	columns, err := MapFuelTypes([]string{"b7"})
	assert.NoError(t, err)
	stations, err := repo.GetStations(ctx, topLeft, bottomRight, []string{}, columns)
	assert.NoError(t, err)
	assert.Len(t, stations, 1)
	assert.Equal(t, "station3", stations[0].NodeID)

	// get all stations in bounding box with brand filter
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandA"}, []string{})
	assert.NoError(t, err)
	assert.Len(t, stations, 2)
	for _, station := range stations {
		assert.Equal(t, "BrandA", station.Brand)
	}

	// get all stations in bounding box with brand and fueltype filter
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandB"}, []string{"e5"})
	assert.NoError(t, err)
	assert.Len(t, stations, 1)
	assert.Equal(t, "BrandB", stations[0].Brand)
	assert.Equal(t, "station2", stations[0].NodeID)

	// upserted stations return their latest details
	stations, err = repo.GetStations(ctx, topLeft, bottomRight, []string{"BrandA"}, []string{})
	assert.NoError(t, err)
	for _, station := range stations {
		if station.NodeID == "station1" {
			assert.Equal(t, "Address 1 Updated", station.Address)
		}
	}

	// unvalidated fuel columns are rejected
	_, err = repo.GetStations(ctx, topLeft, bottomRight, []string{}, []string{"1=1; DROP TABLE stations"})
	assert.Error(t, err)
}

func TestGetStationsLatestPrice(t *testing.T) {
	ctx := context.Background()
	repo, priceRepo, _ := newTestRepos(t)
	writeStations(t, repo)

	older := priceBase
	newer := priceBase.Add(48 * time.Hour)
	err := priceRepo.Insert(ctx, []*dao.StationPrice{
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 139.9, Valid: true},
			EffectiveAt: older,
		},
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 142.9, Valid: true},
			EffectiveAt: newer,
		},
	})
	assert.NoError(t, err)

	topLeft := dao.GeometryPoint{X: -0.5, Y: 51.8}
	bottomRight := dao.GeometryPoint{X: 0.5, Y: 50.0}

	stations, err := repo.GetStations(ctx, topLeft, bottomRight, []string{}, []string{})
	assert.NoError(t, err)

	// each station appears once, with only its newest price attached
	assert.Len(t, stations, 3)
	for _, station := range stations {
		if station.NodeID != "station1" {
			continue
		}
		assert.Equal(t, 142.9, station.E5.Float64)
		assert.True(t, station.EffectiveAt.Valid)
		assert.WithinDuration(t, newer, station.EffectiveAt.Time, time.Second)
	}
}

func TestGetStationsByNodeIDPriceHistory(t *testing.T) {
	ctx := context.Background()
	repo, priceRepo, _ := newTestRepos(t)
	writeStations(t, repo)

	err := priceRepo.Insert(ctx, []*dao.StationPrice{
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 139.9, Valid: true},
			EffectiveAt: priceBase,
		},
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 142.9, Valid: true},
			EffectiveAt: priceBase.Add(24 * time.Hour),
		},
	})
	assert.NoError(t, err)

	// one row per price record, newest first
	history, err := repo.GetStationsByNodeID(ctx, "station1")
	assert.NoError(t, err)
	assert.Len(t, history, 2)
	assert.Equal(t, 142.9, history[0].E5.Float64)
	assert.Equal(t, 139.9, history[1].E5.Float64)
	assert.True(t, history[0].EffectiveAt.Time.After(history[1].EffectiveAt.Time))

	// a station with no prices yet returns a single row with null prices
	history, err = repo.GetStationsByNodeID(ctx, "station2")
	assert.NoError(t, err)
	assert.Len(t, history, 1)
	assert.False(t, history[0].E5.Valid)
	assert.False(t, history[0].EffectiveAt.Valid)
}

func TestMapFuelTypes(t *testing.T) {
	// legacy public names map onto the new price columns
	columns, err := MapFuelTypes([]string{"e5", "e10", "b7", "sdv"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"e5", "e10", "b7_standard", "b7_premium"}, columns)

	// new names pass through
	columns, err = MapFuelTypes([]string{"b7_standard", "b7_premium"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"b7_standard", "b7_premium"}, columns)

	// unknown names are rejected
	_, err = MapFuelTypes([]string{"diesel"})
	assert.Error(t, err)
}

func writeStations(t *testing.T, repo StationRepository) {
	stations := []dao.Station{
		{
			NodeID:   "station1",
			Brand:    "BrandA",
			Address:  "Address 1",
			Postcode: "PC1 1AA",
			Location: dao.GeometryPoint{X: -0.1, Y: 51.5},
		},
		{
			NodeID:   "station2",
			Brand:    "BrandB",
			Address:  "Address 2",
			Postcode: "PC2 2BB",
			Location: dao.GeometryPoint{X: -0.2, Y: 51.6},
		},
		{
			NodeID:   "station3",
			Brand:    "BrandA",
			Address:  "Address 3",
			Postcode: "PC3 3CC",
			Location: dao.GeometryPoint{X: -0.3, Y: 51.7},
		},
		// test same station upserted with new details
		{
			NodeID:   "station1",
			Brand:    "BrandA",
			Address:  "Address 1 Updated",
			Postcode: "PC1 1AA",
			Location: dao.GeometryPoint{X: -0.1, Y: 51.5},
		},
	}

	for _, station := range stations {
		err := repo.Upsert(context.Background(), &station)
		assert.NoError(t, err)
	}
}
