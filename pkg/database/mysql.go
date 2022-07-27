package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(user, pass, db, addr string) (*sql.DB, error) {
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&allowNativePasswords=true", user, pass, addr, db)
	return sql.Open("mysql", dsn)
}
