package service

import (
	"context"
	"main/internal/dao"
	"main/internal/geo"
	"main/internal/model"
	"main/internal/repository"
	"net/http"

	"gorm.io/gorm"
)

const (
	IncorrectCoordsError  = "incorrect number of coordinates provided, must be 4: top-left x, top-left y, bottom-right x, bottom-right y"
	DistanceTooGreatError = "distance between coordinates must be less than 50km"
)

type StationService interface {
	GetBrands(ctx context.Context) ([]string, error)
	GetStations(ctx context.Context, coords []float64, stations []string, fuelTypes []string) ([]model.StationResponse, error)
	GetStationsBySiteID(ctx context.Context, siteID string) ([]model.StationResponse, error)
}

type stationService struct {
	repo repository.StationRepository
}

func NewStationService(db *gorm.DB) StationService {
	return &stationService{
		repo: repository.NewStationRepo(db),
	}
}

// GetBrands retrieves the list of unique station brands. Requires no authentication.
func (s *stationService) GetBrands(ctx context.Context) ([]string, error) {
	res, err := s.repo.GetBrands(ctx)
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    "internal_error",
			Message: "unable to retrieve station brands",
			Status:  http.StatusInternalServerError,
			Err:     err,
		}
	}
	return res, nil
}

// GetStations retrieves stations within the specified bounding box and filters.
func (s *stationService) GetStations(ctx context.Context, coords []float64, stations []string, fuelTypes []string) ([]model.StationResponse, error) {
	if len(coords) != 4 {
		return nil, &model.ErrorResponse{
			Code:    "invalid_input",
			Message: IncorrectCoordsError,
			Status:  http.StatusBadRequest,
		}
	}

	// get points
	topLeft := parseGeometryPoint(coords[:2])
	bottomRight := parseGeometryPoint(coords[2:4])

	// if distance is greater than 50km, return error to prevent expensive query
	distance := geo.CalculateDistance(topLeft, bottomRight)
	if distance > 50000 {
		return nil, &model.ErrorResponse{
			Code:    "invalid_input",
			Message: DistanceTooGreatError,
			Status:  http.StatusBadRequest,
		}
	}

	daoStations, err := s.repo.GetStations(
		ctx,
		topLeft,
		bottomRight,
		stations,
		fuelTypes,
	)

	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    "internal_error",
			Message: "unable to retrieve stations",
			Status:  http.StatusInternalServerError,
			Err:     err,
		}
	}

	var stationResponses []model.StationResponse
	for _, daoStation := range daoStations {
		stationResponses = append(stationResponses, model.FromDAO(&daoStation))
	}

	return stationResponses, nil
}

func (s *stationService) GetStationsBySiteID(ctx context.Context, siteID string) ([]model.StationResponse, error) {
	daoStations, err := s.repo.GetStationsBySiteID(ctx, siteID)
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    "internal_error",
			Message: "unable to retrieve stations with the provided site ID",
			Status:  http.StatusInternalServerError,
			Err:     err,
		}
	}

	if len(daoStations) == 0 {
		return nil, &model.ErrorResponse{
			Code:    "not_found",
			Message: "no stations found with the provided site ID",
			Status:  http.StatusNotFound,
			Err:     gorm.ErrRecordNotFound,
		}
	}

	var stationResponses []model.StationResponse
	for _, daoStation := range daoStations {
		stationResponses = append(stationResponses, model.FromDAO(&daoStation))
	}

	return stationResponses, nil
}

func parseGeometryPoint(coords []float64) dao.GeometryPoint {
	return dao.GeometryPoint{
		X: coords[0],
		Y: coords[1],
	}
}
