package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	// FindTopNCostumers returns the top n customers ordered by money spent.
	TopNCostumers(n int) (c []CustomerMoneySpent, err error)
	// GetTotalMoneySpentByCondition returns the total money spent by a customer.
	GetTotalMoneySpentByCondition(condition int) (float64, error)
}
