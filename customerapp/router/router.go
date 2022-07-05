package router

import (
	"customerapp/controller"

	"github.com/gorilla/mux"
)

func InitRoutes(c *controller.CustomerController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers/{id}", c.Delete).Methods("DELETE")
	r.HandleFunc("/api/customers", c.GetAll).Methods("GET")
	r.HandleFunc("/api/customers/{id}", c.Get).Methods("GET")
	r.HandleFunc("/api/customers", c.Post).Methods("POST")
	r.HandleFunc("/api/customers/{id}", c.Put).Methods("PUT")
	return r
}
