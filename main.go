package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func putData(w http.ResponseWriter, r * http.Request){
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_Long"])
    for k, v := range r.Form{
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
}


func sayRoot(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET"{
    }else{
        putData(w, r)
    }
    t, _ := template.ParseFiles("test.gtpl")
    t.Execute(w, nil)
}

func main(){
    http.HandleFunc("/", sayRoot)
    err := http.ListenAndServe(":9090", nil)
    if err != nil{
        log.Fatal("ListenAndServe: ", err)
    }
        
}
