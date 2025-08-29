package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTableLinks(t *testing.T) {
	url := "https://www.gov.uk/guidance/access-fuel-price-data"
	links, err := GetTableLinks(url)
	if err != nil {
		t.Fatalf("getTableLinks throws error. %v", err)
	}

	assert.NotEmpty(t, links)
}
