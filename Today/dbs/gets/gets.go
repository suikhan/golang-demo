package gets

import (
	"fmt"
	"log"

	"../conns"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers() (string, error) {
	db, err := conns.OpenConn()
	if err != nil {
		return "", err
	}

	rows, err := db.Query("SELECT usercode,username FROM users")
	if err != nil {
		return "", err
	}

	type myrow struct {
		usercode string
		username string
	}

	users := "["
	for rows.Next() {
		var user myrow
		if err := rows.Scan(&user.usercode, &user.username); err != nil {
			log.Fatal("scan错误：", err)
			if err != nil {
				return "", err
			}
		}
		users += "{\"usercode\": \"" + user.usercode + "\",\"username\": \"" + user.username + "\"},"
	}
	users += "]"
	rows.Close()
	return users, err
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
