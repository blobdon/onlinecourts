package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// data for templates
type d map[string]interface{}

// Handler for home page
// /user/:id, i.e. /user/ABC1234567
func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, struct{}{})
}

// Handler for user page
// /user/:id, i.e. /user/ABC1234567
func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	fmt.Println("page", page)
	t, err := template.ParseFiles("static/" + page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, struct{}{})
}

// hangdler for questions
// func questionsHandler(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {

// 	} else if r.Method == http.MethodPost {
// 		http.Redirect(w, r, "/u/")
// 	} else {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 	}

// }

func main() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/user/", mainHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
