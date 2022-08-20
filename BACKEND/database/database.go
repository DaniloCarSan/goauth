package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnection() (db *sql.DB, err error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	db, err = sql.Open("mysql", connStr)
	return
}
