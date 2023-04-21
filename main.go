package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a form Handler that will render the form
func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

// create a handler function to handle the form submission

func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	fmt.Fprintf(w, "Name: %s\nEmail: %s", name, email)
}

// Create a router using Gorilla and define the form
// and submit handlers

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", formHandler).Methods("GET")
	r.HandleFunc("/submit", submitHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
}
