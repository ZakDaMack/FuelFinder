package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StationPrice struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	NodeID string    `gorm:"column:node_id;not null;uniqueIndex:station_prices_node_effective_uniq,priority:1"`

	E5         sql.NullFloat64
	E10        sql.NullFloat64
	B7Standard sql.NullFloat64
	B7Premium  sql.NullFloat64

	EffectiveAt time.Time `gorm:"uniqueIndex:station_prices_node_effective_uniq,priority:2"`
	CreatedAt   time.Time
}

func (s *StationPrice) TableName() string {
	return "station_prices"
}

func (s *StationPrice) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	s.ID = id
	return err
}
