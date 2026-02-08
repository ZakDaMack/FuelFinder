package repository

import (
	"context"
	"fmt"
	"log"
	"main/internal/dao"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StationRepository interface {
	GetBrands(ctx context.Context) ([]string, error)
	Insert(ctx context.Context, station *dao.Station) error
	Exists(ctx context.Context, siteID string, createdAt time.Time) (bool, error)
	GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, includeFueltypes []string) ([]dao.Station, error)
}

type postgresStationRepository struct {
	db *gorm.DB
}

func NewStationRepo(db *gorm.DB) StationRepository {
	if err := db.AutoMigrate(&dao.Station{}); err != nil {
		log.Fatalf("Failed to migrate station schema: %v", err)
	}
	err := db.Exec(`
		CREATE INDEX IF NOT EXISTS stations_location_gix
		ON stations
		USING GIST (location);

		CREATE INDEX IF NOT EXISTS stations_brand_idx
		ON stations (brand);
		
		CREATE UNIQUE INDEX IF NOT EXISTS stations_site_created_uniq
		ON stations (site_id, created_at);
	`).Error
	if err != nil {
		log.Fatalf("Failed to create spatial index on stations location: %v", err)
	}

	return &postgresStationRepository{
		db: db,
	}
}

func (r *postgresStationRepository) GetBrands(ctx context.Context) ([]string, error) {
	var brands []string
	err := r.db.Model(&dao.Station{}).Distinct().Pluck("brand", &brands).Error
	return brands, err
}

func (r *postgresStationRepository) Insert(ctx context.Context, station *dao.Station) error {
	// create state but ignore if site_id and created_at already exist
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "site_id"},
				{Name: "created_at"},
			},
			DoNothing: true,
		}).
		Create(station).Error
}

func (r *postgresStationRepository) GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, includeFueltypes []string) ([]dao.Station, error) {
	var results []dao.Station
	query := r.db.WithContext(ctx).Model(&dao.Station{})
	query = query.Select(`
    	DISTINCT ON (stations.site_id)
		stations.id,
		stations.site_id,
		ST_AsEWKT(stations.location) AS location,
		stations.brand,
		stations.e5,
		stations.e10,
		stations.b7,
		stations.sdv,
		stations.address,
		stations.postcode,
		stations.created_at
	`)

	// Apply bounding box filter
	query = query.Where("stations.location && ST_MakeEnvelope(?, ?, ?, ?, 4326)", topLeft.X, bottomRight.Y, bottomRight.X, topLeft.Y)

	// Apply station filter if provided
	if len(includeBrands) > 0 {
		query = query.Where("stations.brand IN ?", includeBrands)
	}

	// Apply fuel type filter if provided
	if len(includeFueltypes) > 0 {
		for _, fuelType := range includeFueltypes {
			query = query.Where(fmt.Sprintf("stations.%s IS NOT NULL AND stations.%s > 0", fuelType, fuelType))
		}
	}

	// Group stations by site id and select the most recent entry for each station
	query = query.Order("stations.site_id, stations.created_at DESC")

	err := query.Find(&results).Error
	return results, err
}

func (r *postgresStationRepository) Exists(ctx context.Context, siteID string, createdAt time.Time) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&dao.Station{}).
		Where("site_id = ? AND created_at = ?", siteID, createdAt).
		Count(&count).Error
	return count > 0, err
}
