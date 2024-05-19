package conversions

func MilesToRadians(miles float64) float64 {
	return miles / 3963.2
}

func MilesToMetres(miles int) float64 {
	return float64(miles) * 1609.344
}
