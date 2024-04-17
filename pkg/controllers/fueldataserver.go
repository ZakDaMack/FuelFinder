package controllers

import (
	"context"
	"fmt"
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
	writeRes, err := s.store.Write(items.Items)
	if err != nil {
		return nil, err
	}

	fmt.Println(writeRes)
	res := &fueldata.UploadedItems{
		Count: int32(writeRes),
	}

	return res, nil
}
