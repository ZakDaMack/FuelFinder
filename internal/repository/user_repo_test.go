package repository

import (
	"context"
	"main/internal/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCreateUser(t *testing.T) {
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewUserRepo()

	user, err := repo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)
	assert.NotEqual(t, "", user.ID)
	assert.Equal(t, "z.dowsett@email.com", user.Email)
	assert.Equal(t, "Zak", user.FirstName)
	assert.Equal(t, "Dowsett", user.LastName)
}

func TestGetUser(t *testing.T) {
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewUserRepo()
	_, err = repo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	// get user with correct password
	validUser, err := repo.Get(ctx, "z.dowsett@email.com", "password")
	assert.NoError(t, err)
	assert.Equal(t, "z.dowsett@email.com", validUser.Email)

	// try get user with incorrect password
	wrongUser, err := repo.Get(ctx, "z.dowsett@email.com", "wrongpassword")
	assert.Error(t, err)
	assert.Nil(t, wrongUser)

	// try get non-existent user
	nonExistentUser, err := repo.Get(ctx, "wrongguy@email.com", "password")
	assert.Error(t, err)
	assert.Nil(t, nonExistentUser)
}

func TestUpdateUser(t *testing.T) {
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewUserRepo()
	user, err := repo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	user.FirstName = "Zakary"
	user.LastName = "D."
	err = repo.Update(user)
	assert.NoError(t, err)

	updatedUser, err := repo.Get(ctx, "z.dowsett@email.com", "password")
	assert.NoError(t, err)
	assert.Equal(t, "Zakary", updatedUser.FirstName)
	assert.Equal(t, "D.", updatedUser.LastName)
}

func TestDeleteUser(t *testing.T) {
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

	_, err = database.MakePostgresDB(dsn)
	assert.NoError(t, err)

	repo := NewUserRepo()
	NewRefreshTokenRepo() // need to spool migrations for relational integrity
	user, err := repo.Create(ctx, "z.dowsett@email.com", "password", "Zak", "Dowsett")
	assert.NoError(t, err)

	// delete user
	err = repo.Delete(user)
	assert.NoError(t, err)

	// try get deleted user
	deletedUser, err := repo.Get(ctx, "z.dowsett@email.com", "password")
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}
