package database

import (
	"context"
	"log/slog"
	"main/api/fuelfinder"
	"main/internal/conversions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	database string
	client   *mongo.Client
}

type Index struct {
	Key   string
	Value interface{}
}

const (
	_collection = "fuel_data"
)

func NewMongoConnection(uri string, db string) (*MongoStore, error) {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			slog.Debug("mongo started", "command", e.Command.String())
			// fmt.Println(e.Command)
		},
	}

	// TODO: sort out contexts
	options := options.Client().ApplyURI(uri).SetMonitor(monitor)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		return nil, err
	}

	c := &MongoStore{
		database: db,
		client:   client,
	}

	return c, nil
}

func (m *MongoStore) QueryArea(lat, long float64, distanceMiles int, includeBrands []string, includeFueltypes []string) ([]fuelfinder.StationItem, error) {
	coll := m.client.Database(m.database).Collection(_collection)

	distMetres := conversions.MilesToMetres(distanceMiles)
	filter := MakeAggregatePipeline(lat, long, int(distMetres), includeBrands, includeFueltypes)

	// TODO: sort context todo
	cursor, err := coll.Aggregate(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// decode results into array data
	// TODO: is there a mroe effecient way of doing this?
	var results []fuelfinder.StationItem
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// TODO: dont write if the values already exist
func (m *MongoStore) Write(data []*fuelfinder.StationItem) (int, error) {
	docs := make([]interface{}, len(data))
	for i, d := range data {
		docs[i] = &d
	}

	coll := m.client.Database(m.database).Collection(_collection)
	res, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return 0, err
	}
	return len(res.InsertedIDs), nil
}

func (m *MongoStore) Exists(unixTime int64, stationId string) (bool, error) {
	coll := m.client.Database(m.database).Collection(_collection)
	query := bson.D{{Key: "createdat", Value: unixTime}, {Key: "siteid", Value: stationId}}
	err := coll.FindOne(context.TODO(), query).Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	// if both queries above skip, then it must exist
	return true, nil
}

func (m *MongoStore) GetDistinctBrands() ([]string, error) {
	coll := m.client.Database(m.database).Collection(_collection)
	res, err := coll.Distinct(context.TODO(), "brand", bson.D{})
	if err != nil {
		return []string{}, nil
	}

	brands := make([]string, 0)
	for _, b := range res {
		brands = append(brands, b.(string))
	}
	return brands, nil
}

func (m *MongoStore) GetIndexes() (map[string]interface{}, error) {
	coll := m.client.Database(m.database).Collection(_collection)
	res, err := coll.Indexes().List(context.TODO())
	if err != nil {
		return nil, err
	}

	indexes := map[string]interface{}{}
	for res.Next(context.TODO()) {
		var in bson.D
		res.Current.Index(1).Value().Unmarshal(&in)
		indexes[in[0].Key] = in[0].Value
	}
	return indexes, nil
}

func (m *MongoStore) CreateIndex(field string, val interface{}) error {
	coll := m.client.Database(m.database).Collection(_collection)

	indexModel := mongo.IndexModel{Keys: bson.D{{Key: field, Value: val}}}
	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
