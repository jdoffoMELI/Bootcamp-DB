package service

import "app_desafio/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// TopNCostumers returns the top n customers ordered by money spent.
func (s *CustomersDefault) TopNCostumers(n int) (c []internal.CustomerMoneySpent, err error) {
	if n <= 0 {
		return nil, internal.ErrServiceInvalidArgument
	}
	c, err = s.rp.TopNCostumers(n)
	return
}

// GetTotalSalesByCondition returns the total sales by condition.
func (c *CustomersDefault) GetTotalMoneySpentByCondition(condition int) (float64, error) {
	if condition != 0 && condition != 1 {
		return 0, internal.ErrServiceInvalidArgument
	}
	return c.rp.GetTotalMoneySpentByCondition(condition)
}
