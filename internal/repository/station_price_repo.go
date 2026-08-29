package repository

import (
	"context"
	"log"
	"main/internal/dao"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StationPriceRepository interface {
	Insert(ctx context.Context, prices []*dao.StationPrice) error
}

type postgresStationPriceRepository struct {
	db *gorm.DB
}

func NewStationPriceRepo(db *gorm.DB) StationPriceRepository {
	if err := db.AutoMigrate(&dao.StationPrice{}); err != nil {
		log.Fatalf("Failed to migrate station price schema: %v", err)
	}

	return &postgresStationPriceRepository{
		db: db,
	}
}

func (r *postgresStationPriceRepository) Insert(ctx context.Context, prices []*dao.StationPrice) error {
	if len(prices) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "node_id"},
				{Name: "effective_at"},
			},
			DoNothing: true,
		}).
		CreateInBatches(prices, 500).Error
}
