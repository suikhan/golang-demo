package main

import (
  "fmt"
  "net/http"
  "strings"
  "html"
  "io/ioutil"
  //"encoding/json"
)

type Server struct {
  ServerName string
  ServerIP   string
}

type Serverslice struct {
  Servers []Server
  ServersID  string
}

func main() {
  http.HandleFunc("/", handler) 
  http.ListenAndServe(":9002", nil)
}

func handler(w http.ResponseWriter, r *http.Request) { 
  r.ParseForm() 
  fmt.Fprintf(w, "Hi, I love you %s", html.EscapeString(r.URL.Path[1:]))
  if r.Method == "GET" {
    fmt.Println("method:", r.Method) 

    fmt.Println("username", r.Form["username"]) 
    fmt.Println("password", r.Form["password"]) 

    for k, v := range r.Form {
      fmt.Print("key:", k, "; ")
      fmt.Println("val:", strings.Join(v, ""))
    }
  } else if r.Method == "POST" {
    result, _:= ioutil.ReadAll(r.Body)
    r.Body.Close()
    fmt.Printf("%s\n", result)
  } 
}
