package config

// create struct for  mongo config
type MongoConfig struct {
	URL string
	DB  string
}

func GetMongoConfig() *MongoConfig {
	return &MongoConfig{
		URL: "mongodb://localhost:27017",
		DB:  "alert_db",
	}
}
