package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Station struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	SiteID   string
	Location GeometryPoint `gorm:"type:geometry(Point,4326)"`

	E5  sql.NullFloat64 `gorm:"e5"`
	E10 sql.NullFloat64 `gorm:"e10"`
	B7  sql.NullFloat64 `gorm:"b7"`
	SDV sql.NullFloat64 `gorm:"sdv"`

	Brand     string
	Address   string
	Postcode  string
	CreatedAt time.Time
}

func (s *Station) TableName() string {
	return "stations"
}

func (s *Station) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	s.ID = id
	return err
}
