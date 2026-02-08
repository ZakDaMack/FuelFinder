package geo

import (
	"main/internal/dao"
	"math"
)

const (
	earthRadius = 6371e3 // in metres
)

func CalculateDistance(pointA, pointB dao.GeometryPoint) float64 {
	// latitude in radians
	latARad := pointA.X * math.Pi / 180
	latBRad := pointB.X * math.Pi / 180

	// distance in radians
	distanceLatRad := (pointB.X - pointA.X) * math.Pi / 180
	distanceLngRad := (pointB.Y - pointA.Y) * math.Pi / 180

	// calc azimuth
	azimuth := math.Sin(distanceLatRad/2)*math.Sin(distanceLatRad/2) +
		math.Cos(latARad)*math.Cos(latBRad)*
			math.Sin(distanceLngRad/2)*math.Sin(distanceLngRad/2)

	// calc distance
	c := 2 * math.Atan2(math.Sqrt(azimuth), math.Sqrt(1-azimuth))
	return earthRadius * c // in metres
}
