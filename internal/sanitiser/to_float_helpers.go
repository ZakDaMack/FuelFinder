package sanitiser

import "strconv"

func ToFloat(value interface{}) float64 {
	// r := regexp.MustCompile(`("longitude"|"latitude"):(".+?")`)
	// TODO: Fix panic
	var val float64
	switch cast := value.(type) {
	case string:
		val, _ = strconv.ParseFloat(cast, 32)
	case float64:
		val = cast
	}

	return val
}
