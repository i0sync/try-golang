package main

import (
	"customerapp/domain"
	"customerapp/mapstore"
	"log"
)

type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store domain.CustomerStore // CustomerStore value
}

func (cc CustomerController) AddCustomer(c domain.Customer) {
	log.Printf("Inside AddCustomer method")
	err := cc.store.Create(c)
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

func (cc CustomerController) ReadCustomer(Id string) {
	log.Printf("Inside ReadCustomer method")
	customer, err := cc.store.GetById(Id)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println(customer)
}

func (cc CustomerController) ReadAllCustomers() {
	log.Printf("Inside ReadAllCustomers method")
	customers, err := cc.store.GetAll()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	for _, customer := range customers {
		log.Println(customer)
	}
}

func (cc CustomerController) UpdateCustomer(Id string, c domain.Customer) {
	log.Printf("Inside UpdateCustomer method")
	err := cc.store.Update(Id, c)
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

func (cc CustomerController) DeleteCustomer(Id string) {
	log.Printf("Inside DeleteCustomer method")
	err := cc.store.Delete(Id)
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

func main() {
	/*
		Output:
		-------
		Inside AddCustomer method
		New customer [AUNSWGR2003] added !

		Inside ReadCustomer method
		{AUNSWGR2003 Sujith M sujith.m@email.com}

		Inside ReadAllCustomers method
		{AUNSWGR2003 Sujith M sujith.m@email.com}

		Inside UpdateCustomer method
		Customer [AUNSWGR2003] updated !

		Inside ReadAllCustomers method
		{AUNSWGR2003 Sujith M sujith.muraleedharan@email.com}

		Inside DeleteCustomer method
		Error: Customer doesn't exist !

		Inside DeleteCustomer method
		Customer [AUNSWGR2003] deleted !

		Inside ReadAllCustomers method
	*/

	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
		// store : mongodb.NewMongoStore(), // with another database
	}
	c1 := domain.Customer{
		Id:    "AUNSWGR2003",
		Name:  "Sujith M",
		Email: "sujith.m@email.com",
	}

	updated_c1 := domain.Customer{
		Id:    "AUNSWGR2003",
		Name:  "Sujith M",
		Email: "sujith.muraleedharan@email.com",
	}

	controller.AddCustomer(c1)
	controller.ReadCustomer(c1.Id)
	controller.ReadAllCustomers()

	controller.UpdateCustomer(updated_c1.Id, updated_c1)
	controller.ReadAllCustomers()

	// Expected to throw an error
	controller.DeleteCustomer("AUNSWGR2004")

	// Delete existing customer
	controller.DeleteCustomer("AUNSWGR2003")
	controller.ReadAllCustomers()
}
