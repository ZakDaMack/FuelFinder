package repository

import (
	"gorm.io/gorm"
)

const migrationLockID = 20260829

func withMigrationLock(db *gorm.DB, fn func(conn *gorm.DB) error) error {
	return db.Connection(func(conn *gorm.DB) error {
		// pg cant migrate the same table at the same time, lock waits out
		// https://github.com/golang-migrate/migrate/blob/v4.14.1/database/postgres/postgres.go#L169
		if err := conn.Exec("SELECT pg_advisory_lock(?)", migrationLockID).Error; err != nil {
			return err
		}
		defer conn.Exec("SELECT pg_advisory_unlock(?)", migrationLockID)

		return fn(conn)
	})
}
