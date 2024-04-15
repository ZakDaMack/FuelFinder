package database

import (
	"context"
	"main/api/fueldata"
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

func (m *MongoConnection) QueryArea(lat, long float64, distanceMiles int) ([]models.FuelPriceData, error) {
	coll := m.client.Database(m.database).Collection(collection)
	filter := makeAggregatePipeline(lat, long, milesToMetres(float64(distanceMiles)))
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

func (m *MongoConnection) Write(data []fueldata.StationItems) error {
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

func milesToMetres(miles float64) float64 {
	return miles / 1609.344
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

/*
AGGREGATE DATA

[
	{
		$geoNear: {
			key: 'location',
			near: {
				type: 'Point',
				coordinates: [parseFloat(longitude), parseFloat(latitude)]
			},
			distanceField: 'distance',
			maxDistance: milesToMeters(distance),
			spherical: true
		}
	},
	{ $sort: { created_at: -1 } },
	{
		$group: {
			_id: '$site_id',
			records: { $push: '$$ROOT' }
		}
	},
	{
		$replaceRoot: {
			newRoot: { $first: '$records' }
		}
	}
	],

	{ maxTimeMS: 60000, allowDiskUse: true }
*/

func makeAggregatePipeline(lat, long, distMetres float64) bson.A {
	return bson.A{
		bson.D{{
			Key: "$geoNear", Value: bson.D{
				{Key: "key", Value: "location"},
				{Key: "near", Value: bson.D{
					{Key: "type", Value: "Point"},
					{Key: "coordinates", Value: bson.A{long, lat}},
				}},
				{Key: "distanceField", Value: "distance"},
				{Key: "maxDistance", Value: distMetres},
				{Key: "spherical", Value: true},
			},
		}},
		bson.D{{
			Key: "$sort", Value: bson.D{{
				Key: "created_at", Value: -1,
			}},
		}},
		bson.D{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$site_id"},
				{Key: "records", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				},
				}},
		}},
		bson.D{{
			Key: "$replaceRoot", Value: bson.D{{
				Key: "newRoot", Value: bson.D{{
					Key: "$first", Value: "$records",
				}},
			}},
		}},
	}
}
