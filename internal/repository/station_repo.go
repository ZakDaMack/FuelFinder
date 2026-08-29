package repository

import (
	"context"
	"fmt"
	"log"
	"main/internal/dao"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var fuelTypeColumns = map[string]string{
	"e5":          "e5",
	"e10":         "e10",
	"b7":          "b7_standard",
	"b7_standard": "b7_standard",
	"sdv":         "b7_premium",
	"b7_premium":  "b7_premium",
}

var validFuelColumns = map[string]bool{
	"e5":          true,
	"e10":         true,
	"b7_standard": true,
	"b7_premium":  true,
}

func MapFuelTypes(fuelTypes []string) ([]string, error) {
	columns := make([]string, 0, len(fuelTypes))
	for _, fuelType := range fuelTypes {
		column, ok := fuelTypeColumns[fuelType]
		if !ok {
			return nil, fmt.Errorf("unknown fuel type %q", fuelType)
		}
		columns = append(columns, column)
	}
	return columns, nil
}

type StationRepository interface {
	GetBrands(ctx context.Context) ([]string, error)
	Upsert(ctx context.Context, station *dao.Station) error
	GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, fuelColumns []string) ([]dao.StationWithPrice, error)
	GetStationsByNodeID(ctx context.Context, nodeID string) ([]dao.StationWithPrice, error)
}

type postgresStationRepository struct {
	db *gorm.DB
}

func NewStationRepo(db *gorm.DB) StationRepository {
	err := withMigrationLock(db, func(conn *gorm.DB) error {
		if err := conn.AutoMigrate(&dao.Station{}, &dao.StationPrice{}); err != nil {
			return err
		}

		return conn.Exec(`
			CREATE INDEX IF NOT EXISTS stations_location_gix
			ON stations
			USING GIST (location);

			CREATE INDEX IF NOT EXISTS stations_brand_idx
			ON stations (brand);
		`).Error
	})
	if err != nil {
		log.Fatalf("Failed to migrate station schema: %v", err)
	}

	return &postgresStationRepository{
		db: db,
	}
}

func (r *postgresStationRepository) GetBrands(ctx context.Context) ([]string, error) {
	var brands []string
	err := r.db.WithContext(ctx).Model(&dao.Station{}).Distinct("brand").Pluck("brand", &brands).Error
	return brands, err
}

func (r *postgresStationRepository) Upsert(ctx context.Context, station *dao.Station) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "node_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"location", "name", "brand", "address", "postcode", "phone", "updated_at",
			}),
		}).
		Create(station).Error
}

const latestPriceJoin = `
	LEFT JOIN LATERAL (
		SELECT e5, e10, b7_standard, b7_premium, effective_at
		FROM station_prices
		WHERE station_prices.node_id = stations.node_id
		ORDER BY effective_at DESC
		LIMIT 1
	) prices ON TRUE
`

const stationWithPriceColumns = `
	stations.id,
	stations.node_id,
	ST_AsEWKT(stations.location) AS location,
	stations.name,
	stations.brand,
	stations.address,
	stations.postcode,
	stations.phone,
	stations.updated_at,
	prices.e5,
	prices.e10,
	prices.b7_standard,
	prices.b7_premium,
	prices.effective_at
`

func (r *postgresStationRepository) GetStations(ctx context.Context, topLeft, bottomRight dao.GeometryPoint, includeBrands, fuelColumns []string) ([]dao.StationWithPrice, error) {
	query := r.db.WithContext(ctx).Table("stations").
		Select(stationWithPriceColumns).
		Joins(latestPriceJoin)

	// Apply bounding box filter
	query = query.Where("stations.location && ST_MakeEnvelope(?, ?, ?, ?, 4326)", topLeft.X, bottomRight.Y, bottomRight.X, topLeft.Y)

	// Apply brand filter if provided
	if len(includeBrands) > 0 {
		query = query.Where("stations.brand IN ?", includeBrands)
	}

	// Apply fuel type filter if provided
	for _, column := range fuelColumns {
		if !validFuelColumns[column] {
			return nil, fmt.Errorf("unknown fuel column %q", column)
		}
		query = query.Where(fmt.Sprintf("prices.%s IS NOT NULL AND prices.%s > 0", column, column))
	}

	var results []dao.StationWithPrice
	err := query.Order("stations.node_id").Find(&results).Error
	return results, err
}

func (r *postgresStationRepository) GetStationsByNodeID(ctx context.Context, nodeID string) ([]dao.StationWithPrice, error) {
	var results []dao.StationWithPrice
	err := r.db.WithContext(ctx).Table("stations").
		Select(`
			stations.id,
			stations.node_id,
			ST_AsEWKT(stations.location) AS location,
			stations.name,
			stations.brand,
			stations.address,
			stations.postcode,
			stations.phone,
			stations.updated_at,
			prices.e5,
			prices.e10,
			prices.b7_standard,
			prices.b7_premium,
			prices.effective_at
		`).
		Joins(`
			LEFT JOIN station_prices prices
			ON prices.node_id = stations.node_id
		`).
		Where("stations.node_id = ?", nodeID).
		Order("prices.effective_at DESC NULLS LAST").
		Find(&results).Error

	return results, err
}
