package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	tpl, _ = tpl.ParseGlob("home.html")
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)

}
