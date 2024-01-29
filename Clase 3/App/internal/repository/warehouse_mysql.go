package repository

import (
	"database/sql"
	app "db_app/internal"
)

type WarehouseMySQL struct {
	db *sql.DB // MySql connection
}

// CreateNewWarehouseMySQL creates a new WarehouseMySQL
// CreateNewWarehouseMySQL(db *sql.DB) -> (repository *WarehouseMySQL)
// Args:
//		db: MySQL connection
// Returns:
//		reposityory: pointer to WarehouseMySQL

func CreateNewWarehouseMySQL(db *sql.DB) *WarehouseMySQL {
	return &WarehouseMySQL{db: db}
}

// GetWarehouse gets a warehouse by id
// GetWarehouse(id int) -> (warehouse Warehouse, error error)
// Args:
//		id: warehouse id
// Returns:
//		warehouse: Warehouse with the given id
//		error    : error (if exists)

func (w *WarehouseMySQL) GetWarehouse(id int) (app.Warehouse, error) {
	/* Query for the Warehouse */
	queryString := "" +
		"SELECT w.id, w.name, w.address, w.telephone, w.capacity" +
		"  FROM warehouses w " +
		" WHERE w.id = ?   "
	row := w.db.QueryRow(queryString, id)

	if err := row.Err(); err != nil {
		return app.Warehouse{}, err
	}

	/* Scan the Warehouse */
	var warehouse app.Warehouse
	err := row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.Warehouse{}, app.ErrRepositoryWarehouseNotFound
		}
		return app.Warehouse{}, err
	}
	return warehouse, nil
}

// GetWarehouses gets all the warehouses
// GetWarehouses() -> (warehouses []Warehouse, error error)
// Args:
//		none
// Returns:
//		warehouses: slice of Warehouses
//		error     : error (if exists)

func (w *WarehouseMySQL) GetWarehouses() ([]app.Warehouse, error) {
	/* Query for the Warehouse */
	queryString := "" +
		"SELECT w.id, w.name, w.address, w.telephone, w.capacity" +
		"  FROM warehouses w "
	rows, err := w.db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/* Scan the Warehouse */
	var warehouses []app.Warehouse
	for rows.Next() {
		var warehouse app.Warehouse
		err := rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

// CreateWarehouse creates a new warehouse
// CreateWarehouse(ws *app.Warehouse) -> (error error)
// Args:
//		ws: Warehouse to create
// Returns:
//		error: error (if exists)

func (w *WarehouseMySQL) CreateWarehouse(ws *app.Warehouse) error {
	/* Query for the Warehouse */
	queryString := "" +
		"INSERT INTO warehouses (name, address, telephone, capacity) " +
		"VALUES (?, ?, ?, ?) "
	result, err := w.db.Exec(queryString, ws.Name, ws.Address, ws.Telephone, ws.Capacity)
	if err != nil {
		return err
	}

	/* Sets the produt's ID using the new ID provided by the DB */
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ws.ID = int(lastInsertId)

	return nil
}
