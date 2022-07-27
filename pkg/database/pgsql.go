package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPG(user, pass, db, addr string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, addr, db)
	return sql.Open("postgres", dsn)
}
