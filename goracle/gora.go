package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-oci8"
)

func query() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "his_gyyy/hiscss@ORCL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select userno, username from users order by userno")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var usercode, username string
		rows.Scan(&usercode, &username)
		log.Printf("%s\t%s", usercode, username)
	}
	rows.Close()
}

/*
func update() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "username/pwd@ORCL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare(`UPDATE FUB_B set name ='cnm'`)
	result, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	count, _ := result.RowsAffected()
	log.Printf("result count:%d", count)
}
*/
func main() {
	query()
}
