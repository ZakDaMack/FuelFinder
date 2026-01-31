package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	TokenID uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	UserID  uuid.UUID `gorm:"type:uuid"`
	User    User      `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (u *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	u.ID = id
	return err
}
