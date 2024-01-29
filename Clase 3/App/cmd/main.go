package main

import (
	"database/sql"
	"db_app/internal/repository"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "my_db",
	}

	/* Creates the connection to the database */
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/* Checks if the connection is established */
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connected succesfuly!")
	repository := repository.CreateNewProductMySQL(db)
	product, err := repository.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
