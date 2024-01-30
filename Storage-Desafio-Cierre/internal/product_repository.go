package internal

// RepositoryProduct is the interface that wraps the basic methods that a product repository must have.
type RepositoryProduct interface {
	// FindAll returns all products saved in the database.
	FindAll() (p []Product, err error)
	// Save saves a product into the database.
	Save(p *Product) (err error)
	// FindTopNProducts returns the top n products ordered by quantity sold.
	TopNProducts(n int) (p []ProductSaleCount, err error)
}
