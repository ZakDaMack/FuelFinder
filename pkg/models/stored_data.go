package models

import "time"

type FuelPriceData struct {
	CreatedAt time.Time `json:"created_at"`
	SiteId    string    `json:"site_id"`
	Brand     string    `json:"brand"`
	Address   string    `json:"address"`
	Postcode  string    `json:"postcode"`
	Location  struct {
		LocationType string    `json:"type"`
		Coordinates  []float64 `json:"coordinates"` // <field>: [<longitude>, <latitude>]
	} `json:"location"`
	E5  float64 `json:"e5"`
	E10 float64 `json:"e10"`
	B7  float64 `json:"b7"`
	SDV float64 `json:"sdv"`
}
