package internal

type Repository interface {
	GetProduct(id int) (Product, error)
	GetProducts() ([]Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
}

/* TODO: Abstraer los errores de mysql a erroes de repository consistentes con el dominio del problema */
