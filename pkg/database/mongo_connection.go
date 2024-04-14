package database

import (
	"context"
	"main/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	database string
	client   *mongo.Client
}

const (
	collection = "fuel_data"
)

func NewMongoConnection(uri string, db string) (*MongoConnection, error) {
	// TODO: sort out contexts
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		return nil, err
	}

	c := &MongoConnection{
		database: db,
		client:   client,
	}

	return c, nil
}

/*
   const results = await fuelData.find({
       location: {
           $geoWithin: {
               $centerSphere: [
                   [ longitude, latitude ],
                   milesToRadians(distance)
               ]
           }
       }
   }).exec();
*/

func (m *MongoConnection) QueryArea(lat, long float64, distanceMiles int) ([]models.FuelPriceData, error) {
	coll := m.client.Database(m.database).Collection(collection)
	filter := makeFilter(lat, long, milesToRadians(float64(distanceMiles)))
	// get data
	// TODO: sort context todo
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []models.FuelPriceData
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m *MongoConnection) Write(data []models.FuelPriceData) error {
	docs := make([]interface{}, len(data))
	for i, d := range data {
		docs[i] = d
	}

	coll := m.client.Database(m.database).Collection(collection)
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}
	return nil
}

func milesToRadians(miles float64) float64 {
	return miles / 3963.2
}

func makeFilter(lat, long, distRads float64) bson.D {
	return bson.D{{
		Key: "location", Value: bson.D{{
			Key: "$geoWithin", Value: bson.D{{
				Key: "$centerSphere", Value: bson.A{
					bson.A{long, lat},
					distRads,
				},
			}},
		}},
	}}
}
