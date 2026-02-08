package model

import (
	"main/internal/dao"
	"math"
	"time"
)

type StationResponse struct {
	ID        string  `json:"id"`
	SiteID    string  `json:"site_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	E5  *float64 `json:"e5,omitempty"`
	E10 *float64 `json:"e10,omitempty"`
	B7  *float64 `json:"b7,omitempty"`
	SDV *float64 `json:"sdv,omitempty"`

	Brand     string `json:"brand"`
	Address   string `json:"address"`
	Postcode  string `json:"postcode"`
	CreatedAt string `json:"created_at"`
}

func FromDAO(daoStation *dao.Station) StationResponse {
	var e5, e10, b7, sdv *float64

	if daoStation.E5.Valid {
		val := math.Round(daoStation.E5.Float64*10) / 10
		e5 = &val
	}
	if daoStation.E10.Valid {
		val := math.Round(daoStation.E10.Float64*10) / 10
		e10 = &val
	}
	if daoStation.B7.Valid {
		val := math.Round(daoStation.B7.Float64*10) / 10
		b7 = &val
	}
	if daoStation.SDV.Valid {
		val := math.Round(daoStation.SDV.Float64*10) / 10
		sdv = &val
	}

	return StationResponse{
		ID:        daoStation.ID.String(),
		SiteID:    daoStation.SiteID,
		Latitude:  daoStation.Location.Y,
		Longitude: daoStation.Location.X,
		E5:        e5,
		E10:       e10,
		B7:        b7,
		SDV:       sdv,
		Brand:     daoStation.Brand,
		Address:   daoStation.Address,
		Postcode:  daoStation.Postcode,
		CreatedAt: daoStation.CreatedAt.Format(time.RFC3339),
	}
}
