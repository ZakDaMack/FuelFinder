package repository

import (
	"context"
	"main/internal/database"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCreateRefreshToken(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
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

	userRepo := NewUserRepo(db)
	repo := NewRefreshTokenRepo(db)
	user, err := userRepo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	userID := user.ID
	tokenID := uuid.New()

	// create token
	refreshToken, err := repo.Create(ctx, userID, tokenID)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, refreshToken.ID)
	assert.Equal(t, userID, refreshToken.UserID)
	assert.Equal(t, tokenID, refreshToken.TokenID)
}

func TestGetRefreshToken(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
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

	userRepo := NewUserRepo(db)
	repo := NewRefreshTokenRepo(db)
	user, err := userRepo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	userID := user.ID
	tokenID := uuid.New()

	_, err = repo.Create(ctx, userID, tokenID)
	assert.NoError(t, err)

	// get valid token
	retrievedToken, err := repo.Get(ctx, tokenID)
	assert.NoError(t, err)
	assert.Equal(t, userID, retrievedToken.UserID)
	assert.Equal(t, tokenID, retrievedToken.TokenID)

	// get invalid token
	invalidTokenID := uuid.New()
	retrievedToken, err = repo.Get(ctx, invalidTokenID)
	assert.Error(t, err)
	assert.Nil(t, retrievedToken)
}

func TestDeleteRefreshToken(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
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

	userRepo := NewUserRepo(db)
	repo := NewRefreshTokenRepo(db)
	user, err := userRepo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	userID := user.ID
	tokenID := uuid.New()

	res, err := repo.Create(ctx, userID, tokenID)
	assert.NoError(t, err)

	// delete token
	err = repo.Delete(res)
	assert.NoError(t, err)

	// try get deleted token
	deletedToken, err := repo.Get(ctx, tokenID)
	assert.Error(t, err)
	assert.Nil(t, deletedToken)
}

func TestUserRelationCascadeDelete(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
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

	userRepo := NewUserRepo(db)
	repo := NewRefreshTokenRepo(db)
	user, err := userRepo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	userID := user.ID
	tokenID := uuid.New()

	_, err = repo.Create(ctx, userID, tokenID)
	assert.NoError(t, err)

	// delete user
	err = userRepo.Delete(user)
	assert.NoError(t, err)

	// try get token, should be deleted via cascade
	deletedToken, err := repo.Get(ctx, tokenID)
	assert.Error(t, err)
	assert.Nil(t, deletedToken)
}
