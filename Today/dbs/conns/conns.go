package conns

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:shutian1.2@tcp(127.0.0.1:3306)/demo")
	return db, err
}
