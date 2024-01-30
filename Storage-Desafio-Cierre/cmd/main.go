package main

import (
	"app_desafio/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	loadData := false
	if len(os.Args) == 2 {
		if os.Args[1] == "-load" {
			loadData = true
		}
	}

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   "root",
			Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)
	// - set up
	err := app.SetUp(loadData)
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
