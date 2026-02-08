package scraper

import "time"

// incoming data types
type PriceDataset struct {
	LastUpdated string           `json:"last_updated"`
	Stations    []StationDataset `json:"stations"`
}

type StationDataset struct {
	SiteId   string `json:"site_id"`
	Brand    string `json:"brand"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Location struct {
		// FIXME: Is there a better way of handling a mix of string|float64?
		Latitude  any `json:"latitude"`
		Longitude any `json:"longitude"`
	} `json:"location"`
	Prices struct {
		E5  float32 `json:"e5"`
		E10 float32 `json:"e10"`
		B7  float32 `json:"b7"`
		SDV float32 `json:"sdv"`
	} `json:"prices"`
	CreatedAt *time.Time `json:"-"` // to be set when processing
}
