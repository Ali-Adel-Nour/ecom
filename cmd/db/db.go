package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Confing) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
