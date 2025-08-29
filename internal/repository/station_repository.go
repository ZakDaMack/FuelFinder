package repository

import (
	"context"
	"main/api/fuelfinder"
	"main/internal/conversions"
	"main/internal/database"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StationRepository interface {
	GetBrands(ctx context.Context) ([]string, error)
	GetStations(ctx context.Context, lat, long float64, distanceMiles int, includeBrands, includeFueltypes []string) ([]fuelfinder.StationItem, error)
}

type mongoStationRepository struct {
	collection *mongo.Collection
}

func NewMongoStationRepo(db *mongo.Database) StationRepository {
	return &mongoStationRepository{
		collection: db.Collection("fuel_data"),
	}
}

func (r *mongoStationRepository) GetBrands(ctx context.Context) ([]string, error) {
	var brands []string
	res := r.collection.Distinct(ctx, "brand", bson.D{})
	if res.Err() != nil {
		return brands, res.Err()
	}

	if err := res.Decode(&brands); err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *mongoStationRepository) GetStations(ctx context.Context, lat, long float64, distanceMiles int, includeBrands, includeFueltypes []string) ([]fuelfinder.StationItem, error) {
	distMetres := conversions.MilesToMetres(distanceMiles)
	filter := database.MakeAggregatePipeline(lat, long, int(distMetres), includeBrands, includeFueltypes)

	cursor, err := r.collection.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	// decode results into array data
	// TODO: is there a mroe effecient way of doing this?
	var results []fuelfinder.StationItem
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *mongoStationRepository) GetIndexes(ctx context.Context) (map[string]any, error) {
	res, err := r.collection.Indexes().List(ctx)
	if err != nil {
		return nil, err
	}

	indexes := make(map[string]any)
	for res.Next(ctx) {
		var in bson.D
		res.Current.Index(1).Value().Unmarshal(&in)
		indexes[in[0].Key] = in[0].Value
	}
	return indexes, nil
}

func (r *mongoStationRepository) CreateIndex(ctx context.Context, field string, val any) error {
	indexModel := mongo.IndexModel{Keys: bson.D{{Key: field, Value: val}}}
	_, err := r.collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

func (r *mongoStationRepository) Write(ctx context.Context, data []*fuelfinder.StationItem) (int, error) {
	docs := make([]any, len(data))
	for i, d := range data {
		docs[i] = &d
	}

	res, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return 0, err
	}
	return len(res.InsertedIDs), nil
}

func (r *mongoStationRepository) Exists(ctx context.Context, unixTime int64, stationId string) (bool, error) {
	query := bson.D{{Key: "createdat", Value: unixTime}, {Key: "siteid", Value: stationId}}
	err := r.collection.FindOne(ctx, query).Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	// if both queries above skip, then it must exist
	return true, nil
}
