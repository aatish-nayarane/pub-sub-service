package mango

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	
)
//  create singletone connection using  sync.once


var (
	client  *mongo.Client
	once	sync.Once
	connError  error
)
// 1st call  GetMongoConfig for  get uri and db data  
func  GetMongoClient(uri  string) (*mongo.Client, error) {
	once.Do(func() {
		ctx, cancel := 
			context.WithTimeout(
				context.Backgroud(),
				10*time.Second,
			)
		defer cancel()
		client, connectErr = 
			mongo.Connect(
				options.Client().ApplyURI(uri),
			)
	}
	)
	return client, connectErr

}