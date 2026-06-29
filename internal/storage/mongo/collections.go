package mango

import (
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// create mongo multiple  mongo  collecion
type Collection struct {
	Switch *mongo.Collection
}

var (
	instance *Collection
	colOnce  sync.Once
)

//  colname = collection name
func GetCollection(client *mongo.Client, dbName string, colName string) *Collection {
	colOnce.Do(func() {
		db := client.Database(dbName)
		instance = &Collection{
			Switch: db.Collection(colName),
		}
	})
	return instance
}
