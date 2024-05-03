package sanitiser

import (
	"main/api/fueldata"
	"strings"
)

func CleanStationItem(item *fueldata.StationItem) {
	item.Brand = cleanBrandName(item.Brand)
}

func cleanBrandName(s string) string {
	s = strings.Trim(s, " ") // remove any trailing/leading white space
	s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	idx := strings.Index(s, " ")
	if idx != -1 {
		s = s[:idx+1] + strings.ToUpper(string(s[idx+1])) + s[idx+2:]
	}

	return s
}
