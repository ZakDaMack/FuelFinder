package database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDB(ctx context.Context, connectionString, db string) (*mongo.Database, error) {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			slog.Debug("mongo query started", "command", e.Command.String())
			// fmt.Println(e.Command)
		},
	}

	client, err := mongo.Connect(options.Client().ApplyURI(connectionString).SetMonitor(monitor))
	if err != nil {
		return nil, err
	}

	return client.Database(db), nil
}
