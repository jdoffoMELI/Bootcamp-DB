package repository

import (
	"database/sql"

	"app_desafio/internal"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*p).Id = int(id)

	return
}

// TopNProducts returns the top n products ordered by quantity sold.
func (r *ProductsMySQL) TopNProducts(n int) (p []internal.ProductSaleCount, err error) {
	queryString := "" +
		"	 SELECT p.id, p.description, sum(s.quantity) as total" +
		"      FROM products p " +
		"INNER JOIN sales s ON s.product_id = p.id" +
		"  GROUP BY p.id " +
		"  ORDER BY total DESC" +
		"     LIMIT ?"

	res, err := r.db.Query(queryString, n)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var pr internal.ProductSaleCount
		err := res.Scan(&pr.ProductId, &pr.Description, &pr.SaleCount)
		if err != nil {
			return nil, err
		}
		p = append(p, pr)
	}

	return
}
