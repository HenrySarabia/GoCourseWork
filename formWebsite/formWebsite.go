package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	// Add a handler to handle serving static files from a specified directory
	// The reason for using StripPrefix is that you can change the served
	// directory as you please, but keep the reference in HTML the same.
	http.Handle("/ghostDir/", http.StripPrefix("/ghostDir/", http.FileServer(http.Dir("css"))))


	http.HandleFunc("/", root)
	http.HandleFunc("/access", access)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <link rel="stylesheet" href="../ghostDir/home.css">
        <title>Name Check</title>
      </head>
      <body>
      <div>
        <h1>Authorization Check</h1>
        <p>Please enter your name.</p>
        <form action="/access" method="post" accept-charset="utf-8">
	  <input type="text" name="str" placeholder="Enter Name Here..." id="str">
	  <input type="submit" value="Check Now">
        </form>
        </div>
      </body>
    </html>
`

// STEP 1: create a new template - looks like it's automatically created
// STEP 2: parse the string into the template
//  // in lay terms: "give the template your form letter"
//  // in lay terms: "put your form letter into the template"
// STEP 3: execute the template
//  // merge template with data

var grantedTemplate = template.Must(template.New("granted").Parse(grantedTemplateHTML))
var deniedTemplate = template.Must(template.New("denied").Parse(deniedTemplateHTML))

func access(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")
	if strings.EqualFold(strEntered, "Henry") {
		err := grantedTemplate.Execute(w, strEntered)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := deniedTemplate.Execute(w, strEntered)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

const grantedTemplateHTML = `
<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<link rel="stylesheet" href="../ghostDir/granted.css">
			<title>Access Granted</title>
    	</head>
    	<body>
    		<div>
      			<h1><strong>Access Granted</strong></h1>
      			<p>Welcome {{.}}.</p>
      		</div>
    	</body>
  	</html>
`

const deniedTemplateHTML = `
<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<link rel="stylesheet" href="../ghostDir/denied.css">
			<title>Access Denied</title>
		</head>
		<body>
			<div>
				<h1><strong>ACCESS DENIED</strong></h1>
				<p>Sorry, {{.}} is not correct. Try again.</p>
				<form action="/access" method="post" accept-charset="utf-8">
				<input type="text" name="str" placeholder="Enter Name Here..." id="str">
				<input type="submit" value="Check Again">
			</div>
			</form>
		</body>
	</html>
`

// WHY {{html .}}
// http://golang.org/pkg/html/template/#hdr-Typed_Strings

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

/*
http://golang.org/pkg/html/template/
http://www.veracode.com/blog/2013/12/golangs-context-aware-html-templates
HANDLING WEB FORMS
https://cloud.google.com/appengine/docs/go/gettingstarted/handlingforms
*/
