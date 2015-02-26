package main


import (
	"net/http"
	"html/template"
)

type Item struct{
    IntId             int
    StringName        string
    BoolValue         bool
}

var templates = template.Must(template.ParseFiles("./template.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, item *Item) {
	err := templates.ExecuteTemplate(w, tmpl+".html", item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
    item := Item {IntId:200, StringName:"Name 4", BoolValue:true}
	
	renderTemplate(w, "template", &item)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}