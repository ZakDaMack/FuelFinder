package repository

import (
	"context"
	"main/internal/database"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestConcurrentMigration simulates stationapi and the scraper starting at the
// same moment: separate connection pools racing to migrate the same tables.
func TestConcurrentMigration(t *testing.T) {
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

	for round := 0; round < 5; round++ {
		// fresh schema each round so the CREATE TABLEs actually race
		db, err := database.MakePostgresDB(dsn)
		require.NoError(t, err)
		require.NoError(t, db.Exec(`DROP TABLE IF EXISTS station_prices, stations, refresh_tokens, users`).Error)

		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// each goroutine gets its own pool, like a separate process
				pool, err := database.MakePostgresDB(dsn)
				require.NoError(t, err)
				NewStationRepo(pool)
				NewStationPriceRepo(pool)
			}()
		}
		wg.Wait()
	}
}
