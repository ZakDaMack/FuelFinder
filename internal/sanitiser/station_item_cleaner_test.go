package sanitiser

import (
	"main/api/fuelfinder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanStationItem(t *testing.T) {
	item := &fuelfinder.StationItem{
		Brand: "ESSO ",
	}

	CleanStationItem(item)
	assert.Equal(t, "Esso", item.Brand)
}

func TestBrandCapitalisation(t *testing.T) {
	item := &fuelfinder.StationItem{
		Brand: "Apple green ",
	}

	CleanStationItem(item)
	assert.Equal(t, "Apple Green", item.Brand)
}

func TestEmptyBrandname(t *testing.T) {
	item := &fuelfinder.StationItem{
		Brand: "",
	}

	// will panic if test fails
	CleanStationItem(item)
}

// Tests that variations of a prior issue become equivalent
func TestVariantsEquivalent(t *testing.T) {
	itemOne := &fuelfinder.StationItem{Brand: "JET "}
	itemTwo := &fuelfinder.StationItem{Brand: "JET"}
	itemThree := &fuelfinder.StationItem{Brand: "Jet"}

	CleanStationItem(itemOne)
	CleanStationItem(itemTwo)
	CleanStationItem(itemThree)

	if itemOne.Brand != itemTwo.Brand || itemTwo.Brand != itemThree.Brand {
		t.Fatalf("all strings arent equal: %v %v %v", itemOne.Brand, itemTwo.Brand, itemThree.Brand)
	}
}
