package main

import (
    "fmt"
    fb "github.com/huandu/facebook"
)

func main() {
    res, _ := fb.Get("/4", fb.Params{
        "fields": "username",
    })
    fmt.Println("here is my facebook username:", res["username"])
}