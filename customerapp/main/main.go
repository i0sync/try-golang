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

func (c CustomerController) Add(c domain.Customer) {
	err := c.store.Create(c)
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
}
