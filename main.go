package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// data for templates
type d map[string]interface{}

// Handler for home page
// /user/:id, i.e. /user/ABC1234567
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Path[len("/user/"):]
	t, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, d{})
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
		t, err := template.ParseFiles("static/newClaim.html")
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

// handler for exisitng cases
func caseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/case/"):]
	if len(id) == 0 {
		// list all cases
		t, err := template.ParseFiles("static/caseList.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	} else {
		// show details for identified case
		// get id'd case from storage
		// TODO, change this template to case template when completed
		t, err := template.ParseFiles("static/case.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	}
}

// handler for evidence
func evidenceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/evidence.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, d{})
}

func main() {
	db, err := bolt.Open("dummy.db", 0600, nil)
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Println("Connected to db")
	}
	defer db.Close()

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/user/", mainHandler)
	http.HandleFunc("/case/new", newCaseHandler)
	http.HandleFunc("/case/", caseHandler)
	http.HandleFunc("/evidence/", evidenceHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
