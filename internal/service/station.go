package service

import "main/internal/dao"

type StationService interface {
	GetBrands() ([]string, error)
	GetStations(topLeft, bottomRight string, stations []string, fuelTypes []string) ([]dao.Station, error)
}
