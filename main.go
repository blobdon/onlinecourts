package main

import (
	"html/template"
	"net/http"
)

// Handler for home page
func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, struct{}{})
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
