package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email         string    `gorm:"uniqueIndex"`
	FirstName     string
	LastName      string
	Password      string
	RefreshTokens []RefreshToken `gorm:"constraint:OnDelete:CASCADE;"`
	DeletedAt     gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	u.ID = id
	return err
}
