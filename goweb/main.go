package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Anasayfa)
	http.HandleFunc("/detay", Detay)
	http.ListenAndServe(":8080", nil)
}

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "navbar.html")
	view.ExecuteTemplate(w, "anasayfa", nil)
}

func Detay(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detay.html", "navbar.html")
	view.ExecuteTemplate(w, "detay", nil)
}
