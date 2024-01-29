package repository

import (
	"database/sql"
	app "db_app/internal"
)

type ProductMySQL struct {
	db *sql.DB // MySql connection
}

// CreateNewProductMySQL creates a new ProductMySQL
// CreateNewProductMySQL(db *sql.DB) -> (repository *ProductMySQL)
// Args:
//		db: MySQL connection
// Returns:
//		reposityory: pointer to ProductMySQL

func CreateNewProductMySQL(db *sql.DB) *ProductMySQL {
	return &ProductMySQL{db: db}
}

// GetProduct gets a product by id
// GetProduct(id int) -> (product Product, error error)
// Args:
//		id: product id
// Returns:
//		product: Product with the given id
//		error  : error (if exists)

func (p ProductMySQL) GetProduct(id int) (app.Product, error) {
	/* Query for the Product */
	queryString := "SELECT p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price" +
		"  FROM products p " +
		" WHERE p.id = ?   "
	row := p.db.QueryRow(queryString, id)

	if err := row.Err(); err != nil {
		return app.Product{}, err
	}

	/* Scan the Product */
	var product app.Product
	err := row.Scan(&product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return app.Product{}, err
	}

	return product, nil
}

// GetProducts gets all the products
// GetProducts() -> (products []Product, error error)
// Args:
//		none
// Returns:
//		products: slice of Products
//		error   : error (if exists)

func (p ProductMySQL) GetProducts() ([]app.Product, error) {
	/* Query definition */
	queryString := "" +
		"SELECT p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price" +
		"  FROM products p "
	rows, err := p.db.Query(queryString)

	if err != nil {
		return nil, err
	}

	/* Scan the Products */
	var products []app.Product
	for rows.Next() {
		var product app.Product
		err := rows.Scan(&product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// CreateProduct creates a new product
// CreateProduct(product *Product) -> (error error)
// Args:
//		product: pointer to Product
// Returns:
//		error  : error (if exists)

func (p ProductMySQL) CreateProduct(product *app.Product) error {
	/* Query definition */
	queryString := "" +
		"INSERT INTO products (name, quantity, code_value, is_published, expiration, price)" +
		" 	  VALUES (?, ?, ?, ?, ?, ?) "

	result, err := p.db.Exec(
		queryString,
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)

	if err != nil {
		return err
	}

	/* Sets the produt's ID using the new ID provided by the DB */
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.Id = int(lastInsertId)

	return nil
}

// UpdateProduct updates a product
// UpdateProduct(product *Product) -> (error error)
// Args:
//		product: pointer to Product
// Returns:
//		error  : error (if exists)

func (p ProductMySQL) UpdateProduct(product *app.Product) error {
	/* Query definition */
	queryString := "" +
		"UPDATE products p" +
		"   SET p.name = ?, p.quantity = ?, p.code_value = ?, p.is_published = ?, p.expiration = ?, p.price = ?" +
		" WHERE p.id = ? "

	_, err := p.db.Exec(
		queryString,
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
		product.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct deletes a product
// DeleteProduct(id int) -> (error error)
// Args:
//		id: product id
// Returns:
//		error  : error (if exists)

func (p ProductMySQL) DeleteProduct(id int) error {
	/* Query definition */
	queryString := "" +
		"DELETE FROM products p" +
		" WHERE p.id = ?"

	_, err := p.db.Exec(queryString, id)

	if err != nil {
		return err
	}

	return nil
}
