package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "time"
    "crypto/md5"
    "strconv"
    "io"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
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

func writeMessageOnDB(w http.ResponseWriter, r * http.Request, db *sql.DB){
    stmt, err := db.Prepare("INSERT INTO messages(username, message, created) values(?,?,?)")
    checkErr(err, "preparing")


    loc, err := time.LoadLocation("Asia/Tokyo")
    checkErr(err, "loading location")
    username := r.Form.Get("username")
    message := r.Form.Get("message")
    created := time.Now().In(loc)

    res, err := stmt.Exec(username, message, created)
    checkErr(err, "executing")
    _ = res
}

func checkErr(err error, mes string){
    if err != nil{
        println("[message] %v", mes)
        panic(err)
    }
}

func readDB(db *sql.DB) ([](map[string]string)){
    result := make([](map[string]string), 0)
    rows, err := db.Query("SELECT * FROM messages")
    for rows.Next(){
        var uid int
        var username string
        var message string
        var date string
        err = rows.Scan(&uid, &username, &message, &date)
        checkErr(err, "reading")
        result = append(result, map[string]string {"username": username, "message": message, "date": date})
    }
    return result
}


func sayRoot(w http.ResponseWriter, r *http.Request){
    db, err := sql.Open("sqlite3", "./test.db")
    checkErr(err, "reading database")

    
    if r.Method == "GET"{
    }else{
        putData(w, r)
        writeMessageOnDB(w, r, db)

    }
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    params := map[string]interface{} {}
    params["token"] = fmt.Sprintf("%x", h.Sum(nil))

    //row1 := map[string]string{"username": "tarou", "message": "happy!", "date": "2019/1/1"}
    //rows := [](map[string]string){row1}
    rows := readDB(db)
    params["rows"] = rows

    t, _ := template.ParseFiles("test.gtpl")
    t.Execute(w, params)
}

func main(){
    println("hoge")
    http.HandleFunc("/", sayRoot)
    err := http.ListenAndServe(":9090", nil)
    if err != nil{
        log.Fatal("ListenAndServe: ", err)
    }
        
}
