package main

import (
	"html/template"
	"net/http"
)

type userInfo struct {
	Name string
	Email string
	Message string
}

var homeTemplate = template.Must(template.ParseFiles("home.html"))
var formTemplate = template.Must(template.ParseFiles("form.html"))
var linkTemplate = template.Must(template.ParseFiles("link.html"))

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/executeForm", formHandler)
	http.HandleFunc("/link1", link1Handler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	u := userInfo{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	err := formTemplate.Execute(w, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func link1Handler(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("name")+" "+r.FormValue("email")+" "+r.FormValue("message")
	err := linkTemplate.Execute(w, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
