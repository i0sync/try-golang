package controller

import (
	"customerapp/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CustomerController struct {
	// Explicit dependency that hides dependent logic
	Store  domain.CustomerStore // CustomerStore value
	Logger *zap.Logger
}

func (cc *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer cc.Logger.Sync()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := cc.Store.Delete(id); err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("URL", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cc.Logger.Info("Deleted Customer ",
		zap.String("Id ", id),
		zap.String("URL", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

func (cc *CustomerController) Put(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer cc.Logger.Sync()

	vars := mux.Vars(r)
	id := vars["id"]

	var customer domain.Customer
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("URL", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update customer
	if err := cc.Store.Update(id, customer); err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("URL", r.URL.String()),
		)
		if err == domain.CustomerNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cc.Logger.Info("Updated Customer ",
		zap.String("Id ", id),
		zap.String("URL ", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

func (cc *CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer cc.Logger.Sync()
	var customer domain.Customer
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create customer
	if err := cc.Store.Create(customer); err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("URL", r.URL.String()),
		)
		if err == domain.CustomerExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cc.Logger.Info("Created Customer ",
		zap.String("URL ", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

func (cc *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer cc.Logger.Sync()

	vars := mux.Vars(r)
	id := vars["id"]

	// Read customer
	if customer, err := cc.Store.GetById(id); err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("Id", id),
			zap.String("URL", r.URL.String()),
		)
		if err == domain.CustomerNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		j, err := json.Marshal(customer)
		if err != nil {
			cc.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		cc.Logger.Info("Customer ",
			zap.String("Id ", id),
			zap.ByteString("Details ", j),
			zap.String("URL ", r.URL.String()),
		)
		w.Write(j)
		w.WriteHeader(http.StatusOK)
	}
}

func (cc *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Flushing any buffered log entries
	defer cc.Logger.Sync()

	if customers, err := cc.Store.GetAll(); err != nil {
		cc.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == domain.CustomerNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		j, err := json.Marshal(customers)
		if err != nil {
			cc.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		cc.Logger.Info("All Customers ",
			zap.ByteString("Details ", j),
			zap.String("URL", r.URL.String()),
		)
		w.Write(j)
		w.WriteHeader(http.StatusOK)
	}
}
