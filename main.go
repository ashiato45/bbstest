package main

import (
    "fmt"
    //"html/template"
    "log"
    "net/http"
    //"strings"
)

func sayRoot(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello, world!")
}

func main(){
    http.HandleFunc("/", sayRoot)
    err := http.ListenAndServe(":9090", nil)
    if err != nil{
        log.Fatal("ListenAndServe: ", err)
    }
        
}
