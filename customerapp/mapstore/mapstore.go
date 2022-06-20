package mapstore

type MapStore struct {
	// An in-memory store with a map
	// Use Customer.ID as the key of map
	store map[string]domain.Customer 
}

func NewMapStore() *MapStore {
	return &MapStore{ store: make(map[string]domain.Customer)} 
}
