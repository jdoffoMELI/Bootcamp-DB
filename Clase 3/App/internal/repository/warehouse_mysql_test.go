package repository_test

import (
	"database/sql"
	"db_app/internal"
	"db_app/internal/repository"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	/* The test warehouses database by default only has two register */
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "warehouses_test",
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestGetWarehouses(t *testing.T) {
	t.Run("success - should return all the warehouses", func(t *testing.T) {
		/* Initialize the connection */
		db, err := sql.Open("txdb", "warehouses_test")
		require.NoError(t, err)
		defer db.Close()

		/* Initialize the repository */
		repository := repository.CreateNewWarehouseMySQL(db)

		/* Execute the method */
		warehouses, err := repository.GetWarehouses()
		require.NoError(t, err)
		require.NotEmpty(t, warehouses)
		require.Equal(t, 2, len(warehouses))
	})
}

func TestGetWarehouse(t *testing.T) {
	t.Run("success - should return a warehouse", func(t *testing.T) {
		/* Initialize the connection */
		db, err := sql.Open("txdb", "warehouses_test")
		require.NoError(t, err)
		defer db.Close()

		/* Initialize the repository */
		repository := repository.CreateNewWarehouseMySQL(db)

		/* Execute the method */
		warehouse, err := repository.GetWarehouse(1)
		require.NoError(t, err)

		/* Check the result */
		expectedWarehouse := internal.Warehouse{
			ID:        1,
			Name:      "Main Warehouse",
			Address:   "221 Baker Street",
			Telephone: "4555666",
			Capacity:  100,
		}

		require.Equal(t, expectedWarehouse, warehouse)

	})

	t.Run("error - should return an error when the warehouse does not exist", func(t *testing.T) {
		/* Initialize the connection */
		db, err := sql.Open("txdb", "warehouses_test")
		require.NoError(t, err)
		defer db.Close()

		/* Initialize the repository */
		repository := repository.CreateNewWarehouseMySQL(db)

		/* Execute the method */
		warehouse, err := repository.GetWarehouse(3)
		require.Error(t, err)
		require.Equal(t, internal.Warehouse{}, warehouse)
	})
}

func TestCreateWarehouse(t *testing.T) {
	t.Run("success - should create a warehouse", func(t *testing.T) {
		/* Initialize the connection */
		db, err := sql.Open("txdb", "warehouses_test")
		require.NoError(t, err)
		defer db.Close()

		/* Initialize the repository */
		repository := repository.CreateNewWarehouseMySQL(db)

		/* Execute the method */
		warehouse := internal.Warehouse{
			Name:      "Warehouse 3",
			Address:   "221 Baker Street",
			Telephone: "4555666",
			Capacity:  100,
		}
		err = repository.CreateWarehouse(&warehouse)
		warehouses, _ := repository.GetWarehouses()

		/* Check the result */
		require.NoError(t, err)
		require.Equal(t, 3, len(warehouses))
	})
}
