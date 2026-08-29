package model

import (
	"database/sql"
	"main/internal/dao"
	"math"
	"time"
)

type StationResponse struct {
	ID        string  `json:"id"`
	NodeID    string  `json:"node_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	E5         *float64 `json:"e5,omitempty"`
	E10        *float64 `json:"e10,omitempty"`
	B7Standard *float64 `json:"b7_standard,omitempty"`
	B7Premium  *float64 `json:"b7_premium,omitempty"`

	Name     string `json:"name"`
	Brand    string `json:"brand"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone,omitempty"`

	PricesEffectiveAt string `json:"prices_effective_at,omitempty"`
	UpdatedAt         string `json:"updated_at"`
}

func FromDAO(station *dao.StationWithPrice) StationResponse {
	effectiveAt := ""
	if station.EffectiveAt.Valid {
		effectiveAt = station.EffectiveAt.Time.Format(time.RFC3339)
	}

	return StationResponse{
		ID:         station.ID.String(),
		NodeID:     station.NodeID,
		Latitude:   station.Location.Y,
		Longitude:  station.Location.X,
		E5:         roundedPrice(station.E5),
		E10:        roundedPrice(station.E10),
		B7Standard: roundedPrice(station.B7Standard),
		B7Premium:  roundedPrice(station.B7Premium),
		Name:       station.Name,
		Brand:      station.Brand,
		Address:    station.Address,
		Postcode:   station.Postcode,
		Phone:      station.Phone,

		PricesEffectiveAt: effectiveAt,
		UpdatedAt:         station.UpdatedAt.Format(time.RFC3339),
	}
}

func roundedPrice(price sql.NullFloat64) *float64 {
	if !price.Valid {
		return nil
	}

	val := math.Round(price.Float64*10) / 10
	return &val
}
