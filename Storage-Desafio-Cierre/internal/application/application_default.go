package application

import (
	"app_desafio/internal"
	"app_desafio/internal/handler"
	"app_desafio/internal/repository"
	"app_desafio/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationDefault struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Addr is the server address.
	Addr string
}

func loadCustomers(r *repository.CustomersMySQL) {
	/* read de file*/
	f, err := os.Open("./docs/db/json/customers.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/* decoding the data */
	var customers []handler.CustomerJSON
	err = json.NewDecoder(f).Decode(&customers)
	if err != nil {
		panic(err)
	}

	/* deserialization function */
	deserialize := func(c handler.CustomerJSON) internal.Customer {
		attributes := internal.CustomerAttributes{
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}

		return internal.Customer{
			Id:                 c.Id,
			CustomerAttributes: attributes,
		}

	}

	/* load the data */
	for _, customer := range customers {
		c := deserialize(customer)
		err := r.Save(&c)
		if err != nil {
			panic(err)
		}
	}

}

func loadInvoices(r *repository.InvoicesMySQL) {
	/* read de file*/
	f, err := os.Open("./docs/db/json/invoices.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/* decoding the data */
	var invoicesJSON []handler.InvoiceJSON
	err = json.NewDecoder(f).Decode(&invoicesJSON)
	if err != nil {
		panic(err)
	}

	/* deserialization function */
	deserialize := func(i handler.InvoiceJSON) internal.Invoice {
		attributes := internal.InvoiceAttributes{
			Datetime:   i.Datetime,
			Total:      i.Total,
			CustomerId: i.Id,
		}

		return internal.Invoice{
			Id:                i.Id,
			InvoiceAttributes: attributes,
		}

	}

	/* load the data */
	for _, invoice := range invoicesJSON {
		i := deserialize(invoice)
		err := r.Save(&i)
		if err != nil {
			panic(err)
		}
	}
}

func loadProducts(r *repository.ProductsMySQL) {
	/* read de file*/
	f, err := os.Open("./docs/db/json/products.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/* decoding the data */
	var productsJSON []handler.ProductJSON
	err = json.NewDecoder(f).Decode(&productsJSON)
	if err != nil {
		panic(err)
	}

	/* deserialization function */
	deserialize := func(p handler.ProductJSON) internal.Product {
		attributes := internal.ProductAttributes{
			Description: p.Description,
			Price:       p.Price,
		}

		return internal.Product{
			Id:                p.Id,
			ProductAttributes: attributes,
		}

	}

	/* load the data */
	for _, product := range productsJSON {
		p := deserialize(product)
		err := r.Save(&p)
		if err != nil {
			panic(err)
		}
	}
}

func loadSales(r *repository.SalesMySQL) {
	/* read de file*/
	f, err := os.Open("./docs/db/json/sales.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	/* decoding the data */
	var salesJSON []handler.SaleJSON
	err = json.NewDecoder(f).Decode(&salesJSON)
	if err != nil {
		panic(err)
	}

	/* deserialization function */
	deserialize := func(s handler.SaleJSON) internal.Sale {
		attributes := internal.SaleAttributes{
			Quantity:  s.Quantity,
			ProductId: s.ProductId,
			InvoiceId: s.InvoiceId,
		}
		return internal.Sale{
			Id:             s.Id,
			SaleAttributes: attributes,
		}
	}

	/* load the data */
	for _, sale := range salesJSON {
		s := deserialize(sale)
		err := r.Save(&s)
		if err != nil {
			panic(err)
		}
	}
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationDefault{
		cfgDb:   defaultCfg.Db,
		cfgAddr: defaultCfg.Addr,
	}
}

// ApplicationDefault is an implementation of the Application interface.
type ApplicationDefault struct {
	// cfgDb is the database configuration.
	cfgDb *mysql.Config
	// cfgAddr is the server address.
	cfgAddr string
	// db is the database connection.
	db *sql.DB
	// router is the chi router.
	router *chi.Mux
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp(loadData bool) (err error) {
	// dependencies
	// - db: init
	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(a.db)
	rpProduct := repository.NewProductsMySQL(a.db)
	rpInvoice := repository.NewInvoicesMySQL(a.db)
	rpSale := repository.NewSalesMySQL(a.db)
	// - service
	svCustomer := service.NewCustomersDefault(rpCustomer)
	svProduct := service.NewProductsDefault(rpProduct)
	svInvoice := service.NewInvoicesDefault(rpInvoice)
	svSale := service.NewSalesDefault(rpSale)
	// - handler
	hdCustomer := handler.NewCustomersDefault(svCustomer)
	hdProduct := handler.NewProductsDefault(svProduct)
	hdInvoice := handler.NewInvoicesDefault(svInvoice)
	hdSale := handler.NewSalesDefault(svSale)

	if loadData {
		loadCustomers(rpCustomer)
		fmt.Println("Customers loaded")

		loadProducts(rpProduct)
		fmt.Println("Products loaded")

		loadInvoices(rpInvoice)
		fmt.Println("Invoices loaded")

		loadSales(rpSale)
		fmt.Println("Sales loaded")

		fmt.Println("Data loaded sucessfully!")
	}

	// routes
	// - router
	a.router = chi.NewRouter()
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - endpoints
	a.router.Route("/customers", func(r chi.Router) {
		// - GET /customers
		r.Get("/", hdCustomer.GetAll())
		r.Get("/top/{n}", hdCustomer.GetTopNCostumer())
		r.Get("/total/{condition}", hdCustomer.GetTotalMoneySpentByCondition())
		// - POST /customers
		r.Post("/", hdCustomer.Create())
	})
	a.router.Route("/products", func(r chi.Router) {
		// - GET /products
		r.Get("/", hdProduct.GetAll())
		r.Get("/top/{n}", hdProduct.GetTopNProducts())
		// - POST /products
		r.Post("/", hdProduct.Create())
	})
	a.router.Route("/invoices", func(r chi.Router) {
		// - GET /invoices
		r.Get("/", hdInvoice.GetAll())

		// - POST /invoices
		r.Post("/", hdInvoice.Create())
		r.Post("/update", hdInvoice.UpdateTotal())

	})
	a.router.Route("/sales", func(r chi.Router) {
		// - GET /sales
		r.Get("/", hdSale.GetAll())
		// - POST /sales
		r.Post("/", hdSale.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	defer a.db.Close()

	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
