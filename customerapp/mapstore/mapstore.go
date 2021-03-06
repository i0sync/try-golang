package mapstore

import (
	"customerapp/domain"
	"errors"
	"fmt"
)

type MapStore struct {
	// An in-memory store with a map
	// Use Customer.Id as the key of map
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (mstore MapStore) Create(customer domain.Customer) error {
	if _, ok := mstore.store[customer.Id]; ok {
		return errors.New("Customer already exists !")
	}
	mstore.store[customer.Id] = customer
	fmt.Println("New customer added !")
	return nil
}

func (mstore *MapStore) Delete(Id string) error {
	if _, ok := mstore.store[Id]; !ok {
		return errors.New("Customer doesn't exist !")
	}
	delete(mstore.store, Id)
	fmt.Printf("Customer [%s] deleted !", Id)
	return nil
}

func (mstore *MapStore) Update(Id string, customer domain.Customer) error {
	if _, ok := mstore.store[Id]; !ok {
		return errors.New("Customer doesn't exist !")
	}
	mstore.store[customer.Id] = customer
	fmt.Printf("Customer [%s] updated !", Id)
	return nil
}

func (mstore MapStore) GetById(Id string) (domain.Customer, error) {
	var customer domain.Customer
	if cs, ok := mstore.store[Id]; ok {
		return cs, nil
	}
	return customer, errors.New("Customer doesn't exist !")
}

func (mstore MapStore) GetAll() ([]domain.Customer, error) {
	// create a slice of structs to iterate and return the result
	all_customers := make([]domain.Customer, 0, len(mstore.store))
	for _, v := range mstore.store {
		all_customers = append(all_customers, v)
	}
	return all_customers, nil
}
