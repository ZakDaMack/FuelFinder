package scraper

import (
	"log/slog"
	"strings"
)

func IsValidItem(item *StationDataset) bool {
	// if brand name isnt valid, exclude
	if len(item.Brand) == 0 {
		slog.Warn("invalid item", "brand", item.Brand, "id", item.SiteId)
		return false
	}

	return true
}

func cleanStationItem(item *StationDataset) {
	slog.Debug("cleaning item", "brand", item.Brand, "id", item.SiteId)
	item.Brand = cleanBrandName(item.Brand)
}

func cleanBrandName(s string) string {
	if len(s) == 0 {
		return s
	}

	s = strings.Trim(s, " ") // remove any trailing/leading white space
	s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	idx := strings.Index(s, " ") // remove any white space and capitalise the next letter
	if idx != -1 {
		s = s[:idx+1] + strings.ToUpper(string(s[idx+1])) + s[idx+2:]
	}

	return s
}
