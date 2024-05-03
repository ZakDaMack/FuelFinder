package sanitiser

import (
	"main/api/fueldata"
	"main/internal/assert"
	"testing"
)

func TestCleanStationItem(t *testing.T) {
	item := &fueldata.StationItem{
		Brand: "ESSO ",
	}

	CleanStationItem(item)
	assert.Equal("Esso", item.Brand)
}

func TestBrandCapitalisation(t *testing.T) {
	item := &fueldata.StationItem{
		Brand: "Apple green ",
	}

	CleanStationItem(item)
	assert.Equal("Apple Green", item.Brand)
}

// Tests that variations of a prior issue become equivalent
func TestVariantsEquivalent(t *testing.T) {
	itemOne := &fueldata.StationItem{Brand: "JET "}
	itemTwo := &fueldata.StationItem{Brand: "JET"}
	itemThree := &fueldata.StationItem{Brand: "Jet"}

	CleanStationItem(itemOne)
	CleanStationItem(itemTwo)
	CleanStationItem(itemThree)

	if itemOne.Brand != itemTwo.Brand || itemTwo.Brand != itemThree.Brand {
		t.Fatalf("all strings arent equal: %v %v %v", itemOne.Brand, itemTwo.Brand, itemThree.Brand)
	}
}
