package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StationPrice struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Station   Station `gorm:"foreignKey:node_id;references:node_id"`

	E5  sql.NullFloat64 `gorm:"e5"`
	E10 sql.NullFloat64 `gorm:"e10"`
	B7Standard  sql.NullFloat64 `gorm:"b7_standard"`
	B7Premium sql.NullFloat64 `gorm:"b7_premium"`

	EffectiveAt time.Time
	CreatedAt time.Time
}

func (s *Station) TableName() string {
	return "station_prices"
}

func (s *Station) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	s.ID = id
	return err
}
