package main

import (
	"html/template"
	"net/http"
	
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
