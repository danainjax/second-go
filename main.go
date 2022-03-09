package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", SecondServer)
	http.ListenAndServe(":3000", nil)

}

var tpl = template.Must(template.ParseFiles("index.html"))

func SecondServer(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
	w.Write([]byte("<h1> Second Server</h1>"))
}
