package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, page is %s!", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Json page is %s!", r.URL.Path[1:])
}


type Message struct {
    Name string
    Body string
    Time int64
}


func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/json/", jsonHandler)
    http.ListenAndServe(":8080", nil)
}