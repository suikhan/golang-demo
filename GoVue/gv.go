package main

import(
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    _ "strings"
    "io/ioutil"
)
 
func login(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("helloworld.html")
        log.Println(t.Execute(w, nil))
    } else {
        defer r.Body.Close()
        con, _ := ioutil.ReadAll(r.Body)
        fmt.Println(string(con))
        
        var dat map[string]interface{}
        if err:= json.Unmarshal([] byte(string(con)), &dat); err == nil{
            fmt.Println(dat)
            fmt.Println(dat["username"])
        } else {
            fmt.Println(err)
        }
    }
}

func main(){
    h := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", h))
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9077", nil)
    
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
                           
