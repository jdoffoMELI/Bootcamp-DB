package internal

import "errors"

/* Error definitions */
var (
	ErrRepositoryWarehouseNotFound = errors.New("repository: warehouse not found")
)

/* WarehouseRepository interface */
type WarehouseRepository interface {
	GetWarehouse(id int) (Warehouse, error) // Get a warehouse by ID
	GetWarehouses() ([]Warehouse, error)    // Get all the warehouses
	CreateWarehouse(w *Warehouse) error     // Insert a new warehouse

}
