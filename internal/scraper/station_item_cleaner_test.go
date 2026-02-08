package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanStationItem(t *testing.T) {
	item := &StationDataset{
		Brand: "ESSO ",
	}

	cleanStationItem(item)
	assert.Equal(t, "Esso", item.Brand)
}

func TestBrandCapitalisation(t *testing.T) {
	item := &StationDataset{
		Brand: "Apple green ",
	}

	cleanStationItem(item)
	assert.Equal(t, "Apple Green", item.Brand)
}

func TestEmptyBrandname(t *testing.T) {
	item := &StationDataset{
		Brand: "",
	}

	// will panic if test fails
	cleanStationItem(item)
}

// Tests that variations of a prior issue become equivalent
func TestVariantsEquivalent(t *testing.T) {
	itemOne := &StationDataset{Brand: "JET "}
	itemTwo := &StationDataset{Brand: "JET"}
	itemThree := &StationDataset{Brand: "Jet"}

	cleanStationItem(itemOne)
	cleanStationItem(itemTwo)
	cleanStationItem(itemThree)

	if itemOne.Brand != itemTwo.Brand || itemTwo.Brand != itemThree.Brand {
		t.Fatalf("all strings arent equal: %v %v %v", itemOne.Brand, itemTwo.Brand, itemThree.Brand)
	}
}
