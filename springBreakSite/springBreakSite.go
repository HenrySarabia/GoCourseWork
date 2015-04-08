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
	return datastore.NewKey(c, "userInput", "defaultParent", 0, nil)
}

//The initialization function that will call handlers based on the url
func init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/create_message", messageHandler)
}

//This is the main handler that will execute the home page HTML template
func homeHandler(w http.ResponseWriter, r *http.Request){
	//Declaring context and user variables
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

	//Query
	q := datastore.NewQuery("userInput").Ancestor(defaultKey(c)).Order("-Date").Limit(15)
	inputList := make([]userInput, 0, 10)
	if _, err := q.GetAll(c, &inputList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Execute home.html template
	err := homeTemplate.Execute(w, inputList) //Pass in inputList and rewrite HTML
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
		Author: u.String(),
		Content: r.FormValue("messageString"),
		Date: time.Now(),
	}

	key := datastore.NewIncompleteKey(c, "userInput", defaultKey(c))
	_, err := datastore.Put(c, key, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
