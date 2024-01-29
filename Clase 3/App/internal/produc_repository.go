package internal

import "errors"

/* Repository errors */
var (
	ErrRepositoryProductNotFound = errors.New("repository: product not found")
)

/* Interface definition */
type Repository interface {
	GetProduct(id int) (Product, error)   // Get a product by ID
	GetProducts() ([]Product, error)      // Get all products
	CreateProduct(product *Product) error // Insert a new product
	UpdateProduct(product *Product) error // Update a product
	DeleteProduct(id int) error           // Delete a product by ID
}
