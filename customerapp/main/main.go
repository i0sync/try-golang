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

func (cc CustomerController) AddCustomer(c domain.Customer) {
	fmt.Printf("\nInside AddCustomer method\n")
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func (cc CustomerController) ReadCustomer(Id string) {
	fmt.Printf("\nInside ReadCustomer method\n")
	customer, err := cc.store.GetById(Id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(customer)
}

func (cc CustomerController) ReadAllCustomers() {
	fmt.Printf("\nInside ReadAllCustomers method\n")
	customers, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func (cc CustomerController) UpdateCustomer(Id string, c domain.Customer) {
	fmt.Printf("\nInside UpdateCustomer method\n")
	err := cc.store.Update(Id, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func (cc CustomerController) DeleteCustomer(Id string) {
	fmt.Printf("\nInside DeleteCustomer method\n")
	err := cc.store.Delete(Id)
	if err != nil {
		fmt.Println("Error:", err)
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
