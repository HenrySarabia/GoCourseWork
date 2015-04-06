package main

import (
	"net/http"
	"html/template"
	"time"

	"appengine"
	"appengine/user"
	"appengine/datastore"
)

//Defining a struct to hold the user input
type userInput struct {
	Author string
	Content string
	Date time.Time
	}

//Creating an HTML template for the home page
var homeTemplate = template.Must(template.ParseFiles("home.html"))

//Creating a function that, when called, will return a new key with a consistent parent
func defaultKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "userInput", "default", 0, nil)
}

//The initialization function that will call handlers based on the url
func init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/create_message", messageHandler)
}

//This is the main handler that will execute the home page HTML template
func homeHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	w.Header().Set("Location 2", "Sample")
	w.WriteHeader(http.StatusFound)

	err := homeTemplate.Execute(w, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Incomplete. This will be the handler that saves the user's message to the datastore
func messageHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Header", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	w.Header().Set("Header 2", "Sample")
	w.WriteHeader(http.StatusFound)

	input := userInput{
		Content: r.FormValue("")
	}
}

//Still need to add two more handlers and the queries. About 50% done.
