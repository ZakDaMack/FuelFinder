package sanitiser

import (
	"log/slog"
	"main/api/fuelfinder"
	"strings"
)

func IsValidItem(item *fuelfinder.StationItem) bool {
	// if brand name isnt valid, exclude
	if len(item.Brand) == 0 {
		return false
	}

	return true
}

func CleanStationItem(item *fuelfinder.StationItem) {
	slog.Debug("cleaning item", "brand", item.Brand)
	item.Brand = cleanBrandName(item.Brand)
}

func cleanBrandName(s string) string {
	if len(s) == 0 {
		return s
	}

	s = strings.Trim(s, " ") // remove any trailing/leading white space
	s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	idx := strings.Index(s, " ")
	if idx != -1 {
		s = s[:idx+1] + strings.ToUpper(string(s[idx+1])) + s[idx+2:]
	}

	return s
}
