// This is a prototype application created to illustrate what an online
// court experience could be like. The app currently does little more than navigate between
// mocked pages representing different states in the experience for different users.
//
// In order to function as envisioned, we would need to implement many of the
// vital features, a much more complex data model, persistent storage, asynchronous messaging,
// a library of default questions, requirements, and timelines targeted at specific
// types of claims, user authentication, ...

package main

import (
	"html/template"
	"net/http"
)

var dcases, jcases d

// create our mock data
func init() {
	dcases = d{
		"1": d{
			"id":   "HQ17XO1372",
			"c":    "Mrs Connie Wilkinson",
			"d":    "Diamond Care Homes Ltd",
			"s":    "Awaiting Review",
			"dead": "30 Jun 2017",
		},
		"2": d{
			"id":   "FM16P00127",
			"c":    "Mrs Connie Wilkinson",
			"d":    "Mr Fred Bloggs",
			"s":    "Closed",
			"dead": "24 Apr 2015",
		},
	}
	jcases = d{
		"1": d{
			"id":   "HQ17XO1372",
			"c":    "Mrs Connie Wilkinson",
			"d":    "Diamond Care Homes Ltd",
			"s":    "Awaiting Review",
			"dead": "30 Jun 2017",
		},
		"2": d{
			"id":   "452345",
			"c":    "Minerva Hooper",
			"d":    "Cyril Figgis",
			"s":    "Awaiting Defendent Response",
			"dead": "7 Jul 2017",
		},
		"3": d{
			"id":   "897980",
			"c":    "Sage Rodriguez",
			"d":    "Saul Goodman",
			"s":    "Pending Claimant Response",
			"dead": "8 Jul 2017",
		},
		"4": d{
			"id":   "345985",
			"c":    "Philip Chaney",
			"d":    "Sterling Archer",
			"s":    "Pending Claimant Response",
			"dead": "8 Jul 2017",
		},
		"5": d{
			"id":   "767806",
			"c":    "Doris Greene",
			"d":    "Michael Bluth",
			"s":    "Awaiting your Decision",
			"dead": "9 Jul 2017",
		},
	}
}

// data for templates
type d map[string]interface{}

// Handler for home page
// /user/:id, i.e. /user/ABC1234567
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// TODO maintain user authentication/id with token/cookie
	id := r.URL.Path[len("/user/"):]
	// in porduction, here you would get user from db based on id
	// including checking authrntication
	usertype := "party"
	if id == "1" {
		usertype = "judge"
	}
	data := d{
		"jcases":   jcases,
		"usertype": usertype,
	}
	t, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, data)
}

// Handler for html pages
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// TODO maintain user authentication/id with token/cookie
	t, err := template.ParseFiles("static/homepage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, d{})
}

// handler for new case, i.e. /case/new
func newCaseHandler(w http.ResponseWriter, r *http.Request) {
	// TODO maintain user authentication/id with token/cookie
	if r.Method == http.MethodGet {
		// get form to input new case
		t, err := template.ParseFiles("static/newClaim.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	} else if r.Method == http.MethodPost {
		// TODO create/save case
		http.Redirect(w, r, "/case/id#", http.StatusOK)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// handler for exisitng cases
func caseHandler(w http.ResponseWriter, r *http.Request) {
	// TODO maintain user authentication/id with token/cookie
	id := r.URL.Path[len("/case/"):]
	if len(id) == 0 {
		// TODO: get all cases from db with this user ID as a party
		data := d{
			"dcases": dcases,
		}
		t, err := template.ParseFiles("static/caseList.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, data)
	} else {

		t, err := template.ParseFiles("static/case.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, d{})
	}
}

// handler for evidence
func evidenceHandler(w http.ResponseWriter, r *http.Request) {
	// TODO maintain user authentication/id with token/cookie
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
