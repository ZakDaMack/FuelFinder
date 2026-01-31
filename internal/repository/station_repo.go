package repository

import (
	"context"
	"log"
	"main/internal/dao"
	"main/internal/database"

	"gorm.io/gorm"
)

type StationRepository interface {
	GetBrands(ctx context.Context) ([]string, error)
	Insert(ctx context.Context, station *dao.Station) error
	GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, includeFueltypes []string) ([]dao.Station, error)
}

type postgresStationRepository struct {
	db *gorm.DB
}

func NewStationRepo() StationRepository {
	if err := database.DBInstance.AutoMigrate(&dao.Station{}); err != nil {
		log.Fatalf("Failed to migrate station schema: %v", err)
	}
	return &postgresStationRepository{
		db: database.DBInstance,
	}
}

func (r *postgresStationRepository) GetBrands(ctx context.Context) ([]string, error) {
	var brands []string
	err := r.db.Model(&dao.Station{}).Distinct().Pluck("brand", &brands).Error
	return brands, err
}

func (r *postgresStationRepository) Insert(ctx context.Context, station *dao.Station) error {
	return r.db.WithContext(ctx).Save(station).Error
}

func (r *postgresStationRepository) GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, includeFueltypes []string) ([]dao.Station, error) {
	var results []dao.Station
	query := r.db.Model(&dao.Station{})

	// Apply bounding box filter
	query = query.Where("location && ST_MakeEnvelope(?, ?, ?, ?, 4326)", topLeft.X, bottomRight.Y, bottomRight.X, topLeft.Y)

	// Apply station filter if provided
	if len(includeBrands) > 0 {
		query = query.Where("brand IN ?", includeBrands)
	}

	// Apply fuel type filter if provided
	if len(includeFueltypes) > 0 {
		for _, fuelType := range includeFueltypes {
			query = query.Where(fuelType + " IS NOT NULL")
		}
	}

	err := query.Find(&results).Error
	return results, err
}
