package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	// FindTopNCostumers returns the top n customers ordered by money spent.
	TopNCostumers(n int) (c []CustomerMoneySpent, err error)
	// GetTotalMoneySpentByCondition returns the total money spent by a customer.
	GetTotalMoneySpentByCondition(condition int) (float64, error)
}
