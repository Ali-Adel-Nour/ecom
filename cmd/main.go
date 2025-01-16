package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/Ali-Adel-Nour/ecom/cmd/api"
	"github.com/Ali-Adel-Nour/ecom/cmd/config"
	"github.com/Ali-Adel-Nour/ecom/cmd/db"

)

func main() {
	db, err := db.NewMySQLStorage(&mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:              config.Envs.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server starting on http://localhost:3000")
	server := api.NewAPIServer(":3000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	initStorage(db)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
}
