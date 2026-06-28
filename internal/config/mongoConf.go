package  config


// create struct for  mongo config  
type  MongoConfig  struct {
	URL  string
	DB   string
}

func GetMongoConfig() *MongoConfig {
	return &MongoConfig{
		URL    "mongo:localhost:27027"
		DB     "alert_db"
	}
}