package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", r))
}
