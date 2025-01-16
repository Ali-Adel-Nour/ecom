package main

import (
	"fmt"
	"log"

	"github.com/Ali-Adel-Nour/ecom/cmd/api"
	"github.com/Ali-Adel-Nour/ecom/cmd/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "ecom",
		ParseTime:            true,
		AllowNativePasswords: true,
	})

	fmt.Println("Server starting on http://localhost:3000")
	server := api.NewAPIServer(":3000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
