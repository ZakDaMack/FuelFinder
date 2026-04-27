package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJsonFromURL(t *testing.T) {
	// these are problem children, unavailability and antibot measures from the past
	tests := map[string]string{
		"Tesco":  "https://www.tesco.com/fuel_prices/fuel_prices_data.json",
		"Rontec": "https://www.rontec-servicestations.co.uk/fuel-prices/data/fuel_prices_data.json",
		"Esso":   "https://fuelprices.esso.co.uk/latestdata.json",
	}

	for name, url := range tests {
		t.Run(name, func(t *testing.T) {
			links, err := ReadJsonFrom(url)

			assert.NoError(t, err)
			assert.NotEmpty(t, links)
		})
	}
}

func TestHandleDateTime(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:  "valid format 1",
			input: "02/01/2006 15:04:05",
		},
		{
			name:  "valid format 2",
			input: "02-01-2006 15:04:05",
		},
		{
			name:        "invalid format",
			input:       "2006-01-02T15:04:05Z",
			expectError: true,
		},
	}

	// run a subtest for each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := handleDateTime(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
