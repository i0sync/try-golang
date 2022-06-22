package main

import (
	"customerapp/domain"
	"customerapp/mapstore"
	"fmt"
)

type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store domain.CustomerStore // CustomerStore value
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
		// store : mongodb.NewMongoStore(), // with another database
	}
	c1 := domain.Customer{
		Id:    "AUNSWGR2003",
		Name:  "Sujith M",
		Email: "sujith.m@email.com",
	}
	controller.Add(c1)
	controller.GetById("AUNSWGR2003")
}
