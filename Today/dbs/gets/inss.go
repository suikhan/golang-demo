package gets

import (
	"../conns"
	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(userCode string, userName string) {
	db, err := conns.OpenConn()

	stmt, err := db.Prepare(`INSERT users (usercode, username) VALUES (?, ?)`)
	check(err)

	_, err = stmt.Exec(userCode, userName)
	check(err)
}
