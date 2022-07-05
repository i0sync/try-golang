package main

import (
	"customerapp/controller"
	"customerapp/mapstore"
	"customerapp/router"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	ms := mapstore.NewMapStore()

	ctl := controller.CustomerController{
		Store:  ms,
		Logger: logger,
	}

	router := router.InitRoutes(&ctl)
	log.Println("Listening ..")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}
