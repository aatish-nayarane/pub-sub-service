package bootstrap

import (
	"alert_and_notification/internal/config"
	"alert_and_notification/internal/http/handlers"
	"alert_and_notification/internal/services"

	db "alert_and_notification/internal/storage/mongo"
)


func SwitchBootstrap() (*handlers.SwitchHandler, error) {
	// from  cfg you wiil get mongo  url and  and db name 
	cfg:=  config.GetMongoConfig()
	// from client  you will get  mongo  client connection  and  error in  singleton  connection
	client, err := db.GetMongoClient(cfg.URL)
	if err != nil {
		return nil, err
	}
	//  create swithc collection  collection  and  pass  to  service layer
	getSwitchCollection := db.GetCollection(client, cfg.DB, "switch")
	//  setup switch repo
	repo :=  db.NewSwitchRepo(getSwitchCollection)
	//  create  service layer  and pass  collection to service layer

	switchService := services.NewSwitchService(repo)
	//  create handler layer and pass service layer to handler layer
	switchHandler := handlers.NewSwitchHandler(switchService)
	return switchHandler, nil

} 