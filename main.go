package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Case struct {
	ID             string
	ClaimIDs       map[string]string
	AdjudicatorIDs map[string]string
	Status         string
	CreatedAt      string
}

type Claim struct {
	ID           string
	Type         string
	ClaimantIDs  map[string]string
	DefendantIDs map[string]string
	WitnessIDs   map[string]string
	EvidenceIDs  map[string]string
	StdDetails   map[string]string
	SupDetails   map[string]string
	Status       string
	Disposition  string
	CreatedAt    string
}

type Person struct {
	ID         string
	Title      string
	FirstNames string
	Surname    string
	DOB        string
	Address    string
	Postcode   string
	Email      string
	Phone      string
}

type User struct {
	Person
	username string
	passhash string
}

type Evidence struct {
	ID       string
	format   string
	location string // url of stored file
}

type Detail struct {
	ID         string
	QuestionID string
	Response   string
}

type Question struct {
	ID   string
	Text string
}

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
		t, err := template.ParseFiles("static/newclaim.html")
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
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/user/", mainHandler)
	http.HandleFunc("/case/new", newCaseHandler)
	http.HandleFunc("/case/", caseHandler)
	http.HandleFunc("/evidence/", evidenceHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
