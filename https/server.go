package main

import(
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi, this is an example of https service in golang!")
}

func main(){
    http.HandleFunc("/", handler)
    _, err := os.Open("server.crt")
    if err != nil {
        panic(err)
    }
    http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}

