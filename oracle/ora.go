package main

import(
    "database/sql"
    "log"
    _ "github.com/mattn/go-oci8"
)

func main(){
    db, err := sql.Open("oci8", "his_gyyy/hiscss@orcl")
    if err != nil{
        log.Fatal(err)
    }
    rows, err := db.Query("select * from users")
    defer rows.Close()
    if err != nil{
        log.Fatal(err)
    }
    for rows.Next(){
        var i string
        err = rows.Scan(&i)
        if err != nil{
            log.Fatal(err)
        }
        println(i)
    }
}

