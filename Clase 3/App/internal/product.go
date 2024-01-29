package internal

type Product struct {
	Id          int     // Product's ID.
	Name        string  // Product's name.
	Quantity    int     // Product's quantity.
	CodeValue   string  // Product's code value (must be unique).
	IsPublished string  // Tells if the product is published or not.
	Expiration  string  // Product's expiration date in format YYYY-MM-DD.
	Price       float64 // Product's price.
}
