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
	// id :=
	t, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, struct{}{})
}

// Handler for html pages
func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	fmt.Println("page", page)
	t, err := template.ParseFiles("static/" + page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, d{})
}

// handler for new case, i.e. /case/new
func newCaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// get form to input new case
		t, err := template.ParseFiles("static/user.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	} else if r.Method == http.MethodPost {
		// create/save new case
		// TODO create/save case
		http.Redirect(w, r, "/case/id#", http.StatusOK)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// handler for exisitng casese
func caseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/case/"):]
	if len(id) == 0 {
		// list all cases
		t, err := template.ParseFiles("static/table.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	} else {
		// show details for identified case
		// get id'd case from storage
		// TODO, change this template to case template when completed
		t, err := template.ParseFiles("static/user.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	}
}

// handler for evidence
func evidenceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/maps.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, d{})

}

// handler for questions
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
	http.HandleFunc("/case/new", newCaseHandler)
	http.HandleFunc("/case/", caseHandler)
	http.HandleFunc("/evidence/", evidenceHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
