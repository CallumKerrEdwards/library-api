package db

import (
	"context"
	"fmt"
	"os"

	"github.com/CallumKerrEdwards/loggerrific"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client *mongo.Client
	loggerrific.Logger
}

func NewDatabase(ctx context.Context, logger loggerrific.Logger) (*Database, error) {
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}

	return &Database{
		Client: client,
		Logger: logger,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.Ping(ctx, readpref.Primary())
}
