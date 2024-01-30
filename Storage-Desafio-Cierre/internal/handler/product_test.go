package handler_test

import (
	"app_desafio/internal/handler"
	"app_desafio/internal/repository"
	"app_desafio/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestProductsDefault_GetTopProductsByAmountSold tests the handler
func TestTopNProducts(t *testing.T) {
	t.Run("case 1: success - returns top seller products", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		// Roll back routine
		defer func(db *sql.DB) {
			/* Delete records */
			_, err := db.Exec("DELETE FROM sales")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("DELETE FROM products")
			if err != nil {
				panic(err)
			}

			/* Reset auto increment fields */
			_, err = db.Exec("ALTER TABLE sales AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
			if err != nil {
				panic(err)
			}
		}(db)

		// Populate database
		err = func(db *sql.DB) error {
			/* Insert products */
			_, err := db.Exec(
				"INSERT INTO products (`id`, `description`, `price`) VALUES" +
					"(1, 'product 1', 10.00)," +
					"(2, 'product 2', 20.00)," +
					"(3, 'product 3', 30.00)," +
					"(4, 'product 4', 40.00)," +
					"(5, 'product 5', 50.00)," +
					"(6, 'product 6', 60.00);",
			)
			if err != nil {
				return err
			}

			/* Inser sales */
			_, err = db.Exec(
				"INSERT INTO sales (`id`, `product_id`, `quantity`) VALUES" +
					"(1, 1, 500)," +
					"(2, 2, 400)," +
					"(3, 3, 300)," +
					"(4, 4, 200)," +
					"(5, 5, 100);",
			)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		/* Instantiate handler */
		rp := repository.NewProductsMySQL(db) // repository
		sv := service.NewProductsDefault(rp)  // service
		hd := handler.NewProductsDefault(sv)  // handler
		hdFunc := hd.GetTopNProducts()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/products/top/", nil)
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
					"description": "product 1",
					"sale_count": 500
				},
				{
					"description": "product 2",
					"sale_count": 400
				},
				{
					"description": "product 3",
					"sale_count": 300
				},
				{
					"description": "product 4",
					"sale_count": 200
				},
				{
					"description": "product 5",
					"sale_count": 100
				}
			]
		}
		`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
	})
	t.Run("case 2: success - returns no products", func(t *testing.T) {
		/* Preparing dependencies */
		// Database connection
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		/* Instantiate handler */
		rp := repository.NewProductsMySQL(db) // repository
		sv := service.NewProductsDefault(rp)  // service
		hd := handler.NewProductsDefault(sv)  // handler
		hdFunc := hd.GetTopNProducts()

		/* Request definition */
		request := httptest.NewRequest(http.MethodGet, "/products/top", nil)
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
