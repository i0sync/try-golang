package domain

import (
	"errors"
)

var CustomerNotFound = errors.New("No customer found")
var CustomerExists = errors.New("Customer exists")

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerStore interface {
	Create(Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
	Update(string, Customer) error
}
