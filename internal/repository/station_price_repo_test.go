package repository

import (
	"context"
	"database/sql"
	"main/internal/dao"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertPrices(t *testing.T) {
	ctx := context.Background()
	stationRepo, priceRepo, db := newTestRepos(t)
	writeStations(t, stationRepo)

	// an empty batch is a no-op
	err := priceRepo.Insert(ctx, nil)
	assert.NoError(t, err)

	prices := []*dao.StationPrice{
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 139.9, Valid: true},
			E10:         sql.NullFloat64{Float64: 134.5, Valid: true},
			EffectiveAt: priceBase,
		},
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 141.2, Valid: true},
			EffectiveAt: priceBase.Add(24 * time.Hour),
		},
	}

	err = priceRepo.Insert(ctx, prices)
	assert.NoError(t, err)

	var results []dao.StationPrice
	err = db.Model(&dao.StationPrice{}).
		Where("node_id = ?", "station1").
		Order("effective_at ASC").
		Find(&results).Error
	assert.NoError(t, err)
	assert.Len(t, results, 2)

	// values round-trip, including nulls for unreported fuels
	assert.Equal(t, 139.9, results[0].E5.Float64)
	assert.Equal(t, 134.5, results[0].E10.Float64)
	assert.False(t, results[0].B7Standard.Valid)
	assert.False(t, results[0].B7Premium.Valid)
	assert.Equal(t, 141.2, results[1].E5.Float64)
	assert.False(t, results[1].E10.Valid)
}

func TestInsertPricesDuplicate(t *testing.T) {
	ctx := context.Background()
	stationRepo, priceRepo, db := newTestRepos(t)
	writeStations(t, stationRepo)

	price := dao.StationPrice{
		NodeID:      "station1",
		E5:          sql.NullFloat64{Float64: 139.9, Valid: true},
		EffectiveAt: priceBase,
	}
	err := priceRepo.Insert(ctx, []*dao.StationPrice{&price})
	assert.NoError(t, err)

	// same node and effective timestamp again with a different value
	dupe := dao.StationPrice{
		NodeID:      "station1",
		E5:          sql.NullFloat64{Float64: 999.9, Valid: true},
		EffectiveAt: priceBase,
	}
	err = priceRepo.Insert(ctx, []*dao.StationPrice{&dupe})
	assert.NoError(t, err) // Should not error due to ON CONFLICT DO NOTHING

	// only the original record survives
	var results []dao.StationPrice
	err = db.Model(&dao.StationPrice{}).Where("node_id = ?", "station1").Find(&results).Error
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, 139.9, results[0].E5.Float64)
}

func writePrices(t *testing.T, repo StationPriceRepository) {
	prices := []*dao.StationPrice{
		{
			NodeID:      "station1",
			E5:          sql.NullFloat64{Float64: 139.9, Valid: true},
			E10:         sql.NullFloat64{Float64: 134.5, Valid: true},
			EffectiveAt: priceBase,
		},
		{
			NodeID:      "station2",
			E5:          sql.NullFloat64{Float64: 141.0, Valid: true},
			EffectiveAt: priceBase,
		},
		{
			NodeID:      "station3",
			B7Standard:  sql.NullFloat64{Float64: 149.9, Valid: true},
			B7Premium:   sql.NullFloat64{Float64: 155.0, Valid: true},
			EffectiveAt: priceBase,
		},
	}

	err := repo.Insert(context.Background(), prices)
	assert.NoError(t, err)
}
