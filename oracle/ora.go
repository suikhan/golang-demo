package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-oci8"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("OracleDriver example")
	os.Setenv("NLS_LANG", "")

	db, err := sql.Open("oci8", "his_gyyy/hiscss@orcl")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			log.Fatal(err)
		}
		println(i)
	}
	rows.Close()
	db.Close()
}
