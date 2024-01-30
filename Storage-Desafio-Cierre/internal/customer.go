package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerMoneySpend is the struct that represents the money spend by a customer for top N customers.
type CustomerMoneySpent struct {
	CustomerId int     // Id of the customer.
	FirstName  string  // First name of the customer.
	LastName   string  // Last name of the customer.
	TotalSpent float64 // Money spent by the customer.
}
