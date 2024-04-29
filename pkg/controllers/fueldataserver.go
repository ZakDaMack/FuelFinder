package controllers

import (
	"context"
	"log/slog"
	"main/api/fueldata"
	"main/pkg/services"
)

type FuelDataServer struct {
	store *services.MongoStore
	fueldata.UnimplementedFuelDataServer
}

func NewFuelDataServer(db, uri string) (*FuelDataServer, error) {
	conn, err := services.NewMongoConnection(uri, db)
	if err != nil {
		return nil, err
	}

	fds := &FuelDataServer{store: conn}
	return fds, nil
}

func (s *FuelDataServer) QueryArea(ctx context.Context, fence *fueldata.Geofence) (*fueldata.StationItems, error) {
	queryRes, err := s.store.QueryArea(float64(fence.Latitude), float64(fence.Longitude), int(fence.Radius))
	if err != nil {
		return nil, err
	}

	// REVIEW - is there a better way for this?
	// map to pointers
	items := make([]*fueldata.StationItem, len(queryRes))
	for i := 0; i < len(queryRes); i++ {
		items[i] = &queryRes[i]
	}

	results := &fueldata.StationItems{Items: items}
	return results, nil
}

func (s *FuelDataServer) Upload(ctx context.Context, items *fueldata.StationItems) (*fueldata.UploadedItems, error) {
	res := &fueldata.UploadedItems{}

	// check if files already exists, if so, return
	exists, err := s.store.Exists(items.Items[0].CreatedAt, items.Items[0].SiteId)
	if err != nil {
		return nil, err
	}

	if exists {
		res.Count = 0
		return res, nil
	}

	writeRes, err := s.store.Write(items.Items)
	if err != nil {
		return nil, err
	}

	slog.Debug("uploaded station data", "stations", writeRes)
	res.Count = int32(writeRes)

	return res, nil
}

// func (s *FuelDataServer) DistinctBrands(ctx context.Context) (*fueldata.Brands, error) {
// 	brands, err := s.store.GetDistinctBrands()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return brands, nil
// }
