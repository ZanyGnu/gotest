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

type PageDataContext struct{
    Title	string
    Items   []Item
}

var templates = template.Must(template.ParseFiles("./template.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, context *PageDataContext) {
	err := templates.ExecuteTemplate(w, tmpl+".html", context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
    item := Item {IntId:1, StringName:"Name 1", BoolValue:true}
	item2 := Item {IntId:2, StringName:"Name 2", BoolValue:true}

	slice := []Item{}
	slice = append(slice, item)
	slice = append(slice, item2)

	PageDataContext := PageDataContext { Title:"Page Title", Items:slice }

	renderTemplate(w, "template", &PageDataContext)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}