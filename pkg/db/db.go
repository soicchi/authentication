package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() (err error) {
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBName)
	db, err = sql.Open("mysql", connectInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
