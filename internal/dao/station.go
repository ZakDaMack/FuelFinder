package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Station struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	NodeID   string `gorm:"node_id"`
	Location GeometryPoint `gorm:"type:geometry(Point,4326)"`

	Name 	  string
	Brand     string
	Address   string
	Postcode  string
	Phone	  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Station) TableName() string {
	return "stations"
}

func (s *Station) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	s.ID = id
	return err
}
