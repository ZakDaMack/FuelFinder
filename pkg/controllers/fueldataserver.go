package controllers

import (
	"context"
	"main/api/fueldata"
)

type FuelDataServer struct {
	fueldata.UnimplementedFuelDataServer
}

// mustEmbedUnimplementedFuelDataServer implements fueldata.FuelDataServer.
// func (s *FuelDataServer) mustEmbedUnimplementedFuelDataServer() {
// 	panic("unimplemented")
// }

func (s *FuelDataServer) QueryArea(context.Context, *fueldata.Geofence) (*fueldata.StationItems, error) {
	// coll := m.client.Database(m.database).Collection(collection)
	// filter := makeFilter(lat, long, milesToRadians(float64(distanceMiles)))
	// // get data
	// // TODO: sort context todo
	// cursor, err := coll.Find(context.TODO(), filter)
	// if err != nil {
	// 	return nil, err
	// }

	// var results []models.FuelPriceData
	// err = cursor.All(context.TODO(), &results)
	// if err != nil {
	// 	return nil, err
	// }

	item := &fueldata.StationItem{}
	results := &fueldata.StationItems{
		Items: []*fueldata.StationItem{item},
	}

	return results, nil
}

func (s *FuelDataServer) Upload(context.Context, *fueldata.StationItems) (*fueldata.UploadedItems, error) {
	// docs := make([]interface{}, len(data))
	// for i, d := range data {
	// 	docs[i] = d
	// }

	// coll := m.client.Database(m.database).Collection(collection)
	// _, err := coll.InsertMany(context.TODO(), docs)
	// if err != nil {
	// 	return err
	// }

	items := &fueldata.UploadedItems{
		Count: 10,
	}

	return items, nil
}
