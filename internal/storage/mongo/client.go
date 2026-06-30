package mango

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// create singleton connection using sync.Once

var (
	client    *mongo.Client
	once      sync.Once
	connError error
)

// GetMongoClient creates Mongo connection once and reuses it
func GetMongoClient(
	uri string,
) (*mongo.Client, error) {

	once.Do(func() {

		ctx, cancel :=
			context.WithTimeout(
				context.Background(),
				10*time.Second,
			)

		defer cancel()

		client, connError =
			mongo.Connect(
				options.Client().
					ApplyURI(uri),
			)

		if connError != nil {
			return
		}

		// optional ping check
		connError =
			client.Ping(
				ctx,
				nil,
			)
	})

	return client, connError
}