package repository

import (
	"database/sql"

	"app_desafio/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

// TopNProducts returns the top n products ordered by quantity sold.
func (cmsql *CustomersMySQL) TopNCostumers(n int) (c []internal.CustomerMoneySpent, err error) {
	queryString := "" +
		"	 SELECT c.id, c.first_name, c.last_name, SUM(i.total) as total_spent" +
		"	   FROM customers c " +
		"INNER JOIN invoices i ON i.customer_id = c.id" +
		" 	  WHERE c.condition = 1 " +
		"  GROUP BY c.id " +
		"  ORDER BY total_spent DESC" +
		"     LIMIT ?"

	res, err := cmsql.db.Query(queryString, n)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var cms internal.CustomerMoneySpent
		err := res.Scan(&cms.CustomerId, &cms.FirstName, &cms.LastName, &cms.TotalSpent)
		if err != nil {
			return nil, err
		}
		c = append(c, cms)
	}

	return
}

// GetTotalSalesByCondition returns the total sales by a condition.
func (c *CustomersMySQL) GetTotalMoneySpentByCondition(condition int) (float64, error) {
	queryString := "" +
		"    SELECT ROUND(SUM(i.total), 2) as total " +
		"	   FROM customers c " +
		"INNER JOIN invoices i ON c.id = i.customer_id " +
		"	  WHERE c.`condition` = ? " +
		"  GROUP BY c.`condition` "
	row := c.db.QueryRow(queryString, condition)
	var total float64
	err := row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
