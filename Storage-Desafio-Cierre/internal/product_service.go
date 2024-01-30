package internal

import "errors"

var (
	ErrServiceInvalidArgument = errors.New("service: invalid argument")
)

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)
	// Save saves a product.
	Save(p *Product) (err error)
	// FindTopNProducts returns the top n products ordered by quantity sold.
	TopNProducts(n int) (p []ProductSaleCount, err error)
}
