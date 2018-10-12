package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:shutian1.2@tcp(127.0.0.1:3306)/demo")
	checkErr(err)
	stmt, err := db.Prepare("insert demo(name) values(?)")
	checkErr(err)
	res, err := stmt.Exec("ZhangSan")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	fmt.Println("It is over!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
