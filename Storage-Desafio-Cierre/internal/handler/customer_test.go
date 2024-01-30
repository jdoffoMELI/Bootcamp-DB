package handler_test

import (
	"app_desafio/internal/handler"
	"app_desafio/internal/repository"
	"app_desafio/internal/service"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// init registers txdb
func init() {
	// db connection
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Addr:   "127.0.0.1:3306",
		Net:    "tcp",
		DBName: "fantasy_products_test_db",
	}
	// register txdb
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

// addURLParams adds the URL params to the request (needed by Chi framework)
// addURLParams(*http.Request, map[string]string) -> *http.Request
// Args:
// 	req: Request
// 	params: URL params
// Returns:
// 	*http.Request: Request with the URL params

func addURLParams(req *http.Request, params map[string]string) *http.Request {
	chiCtx := chi.NewRouteContext()
	for key, value := range params {
		chiCtx.URLParams.Add(key, value)
	}
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
}

func TestGetTopNCostumer(t *testing.T) {
	t.Run("case 1: success - returns top active customers by amount spent", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		// Roll back routine
		defer func(db *sql.DB) {
			/* Delete records */
			_, err := db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM customers")
			if err != nil {
				panic(err)
			}

			/* Reset auto increment fields */
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)

		// Populate database
		err = func(db *sql.DB) error {
			/* Insert consumers */
			_, err := db.Exec(
				"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES " +
					"(1, 'John', 'Doe', 1), " +
					"(2, 'Jane', 'Doe', 1), " +
					"(3, 'John', 'Smith', 1), " +
					"(4, 'Jane', 'Smith', 1), " +
					"(5, 'John', 'Clark', 1), " +
					"(6, 'Jane', 'Clark', 1), " +
					"(7, 'John', 'Else', 0), " +
					"(8, 'Jane', 'Else', 0);",
			)
			if err != nil {
				return err
			}

			/* Inser invoices */
			_, err = db.Exec(
				"INSERT INTO invoices (`id`, `customer_id`, `total`) VALUES " +
					"(1, 1, 1000), " +
					"(2, 2, 500), " +
					"(3, 3, 250), " +
					"(4, 4, 125), " +
					"(5, 5, 50), " +
					"(6, 6, 25), " +
					"(7, 7, 10), " +
					"(8, 8, 5);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		/* Instantiate handler */
		rp := repository.NewCustomersMySQL(db) // repository
		sv := service.NewCustomersDefault(rp)  // service
		hd := handler.NewCustomersDefault(sv)  // handler
		hdFunc := hd.GetTopNCostumer()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/customers/top", nil)
		request = addURLParams(request, map[string]string{"n": "5"})

		/* Response definition */
		response := httptest.NewRecorder()
		hdFunc(response, request)

		/* Assertions */
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"data": [
					{
						"first_name": "John",
						"last_name": "Doe",
						"total_spent": 1000
					},
					{
						"first_name": "Jane",
						"last_name": "Doe",
						"total_spent": 500
					},
					{
						"first_name": "John",
						"last_name": "Smith",
						"total_spent": 250
					},
					{
						"first_name": "Jane",
						"last_name": "Smith",
						"total_spent": 125
					},
					{
						"first_name": "John",
						"last_name": "Clark",
						"total_spent": 50
					}
				]
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("case 2: success - returns no customers", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		/* Instantiate handler */
		rp := repository.NewCustomersMySQL(db) // repository
		sv := service.NewCustomersDefault(rp)  // service
		hd := handler.NewCustomersDefault(sv)  // handler
		hdFunc := hd.GetTopNCostumer()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/customers/top", nil)
		request = addURLParams(request, map[string]string{"n": "5"})

		/* Request execution */
		response := httptest.NewRecorder()
		hdFunc(response, request)

		/* Assertions */
		expectedCode := http.StatusOK
		expectedBody := `{"data": []}`

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}

func TestGetTotalMoneySpentByCondition(t *testing.T) {
	t.Run("case 1: success - returns total money spent by condition active", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		// Roll back routine
		defer func(db *sql.DB) {
			/* Delete records */
			_, err := db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM customers")
			if err != nil {
				panic(err)
			}

			/* Reset auto increment fields */
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)

		// Populate database
		err = func(db *sql.DB) error {
			/* Insert consumers */
			_, err := db.Exec(
				"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES " +
					"(1, 'John', 'Doe', 1), " +
					"(2, 'Jane', 'Doe', 1), " +
					"(3, 'John', 'Smith', 1), " +
					"(4, 'Jane', 'Smith', 1), " +
					"(5, 'John', 'Clark', 1), " +
					"(6, 'Jane', 'Clark', 1), " +
					"(7, 'John', 'Else', 0), " +
					"(8, 'Jane', 'Else', 0);",
			)
			if err != nil {
				return err
			}

			/* Inser invoices */
			_, err = db.Exec(
				"INSERT INTO invoices (`id`, `customer_id`, `total`) VALUES " +
					"(1, 1, 1000), " +
					"(2, 2, 500), " +
					"(3, 3, 200), " +
					"(4, 4, 100), " +
					"(5, 5, 50), " +
					"(6, 6, 25), " +
					"(7, 7, 10), " +
					"(8, 8, 5);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		/* Instantiate handler */
		rp := repository.NewCustomersMySQL(db) // repository
		sv := service.NewCustomersDefault(rp)  // service
		hd := handler.NewCustomersDefault(sv)  // handler
		hdFunc := hd.GetTotalMoneySpentByCondition()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/customers/top", nil)
		request = addURLParams(request, map[string]string{"condition": "1"})

		/* Response definition */
		response := httptest.NewRecorder()
		hdFunc(response, request)

		/* Assertions */
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"condition": "Activo",
				"total": 1875.00
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})

	t.Run("case 2: success - returns total money spent by condition inactive", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		// Roll back routine
		defer func(db *sql.DB) {
			/* Delete records */
			_, err := db.Exec("DELETE FROM invoices")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM customers")
			if err != nil {
				panic(err)
			}

			/* Reset auto increment fields */
			_, err = db.Exec("ALTER TABLE invoices AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)

		// Populate database
		err = func(db *sql.DB) error {
			/* Insert consumers */
			_, err := db.Exec(
				"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES " +
					"(1, 'John', 'Doe', 1), " +
					"(2, 'Jane', 'Doe', 1), " +
					"(3, 'John', 'Smith', 1), " +
					"(4, 'Jane', 'Smith', 1), " +
					"(5, 'John', 'Clark', 1), " +
					"(6, 'Jane', 'Clark', 1), " +
					"(7, 'John', 'Else', 0), " +
					"(8, 'Jane', 'Else', 0);",
			)
			if err != nil {
				return err
			}

			/* Inser invoices */
			_, err = db.Exec(
				"INSERT INTO invoices (`id`, `customer_id`, `total`) VALUES " +
					"(1, 1, 1000), " +
					"(2, 2, 500), " +
					"(3, 3, 200), " +
					"(4, 4, 100), " +
					"(5, 5, 50), " +
					"(6, 6, 25), " +
					"(7, 7, 10), " +
					"(8, 8, 5);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		/* Instantiate handler */
		rp := repository.NewCustomersMySQL(db) // repository
		sv := service.NewCustomersDefault(rp)  // service
		hd := handler.NewCustomersDefault(sv)  // handler
		hdFunc := hd.GetTotalMoneySpentByCondition()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/customers/top", nil)
		request = addURLParams(request, map[string]string{"condition": "0"})

		/* Response definition */
		response := httptest.NewRecorder()
		hdFunc(response, request)

		/* Assertions */
		expectedCode := http.StatusOK
		expectedBody := `
			{
				"condition": "Inactivo",
				"total": 15.00
			}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
}
