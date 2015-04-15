package main

import (
	"html/template"
	"net/http"
	"time"
	"log"
	"encoding/json"
)

var homeTemplate = template.Must(template.ParseFiles("home.html"))
var formTemplate = template.Must(template.ParseFiles("form.html"))
var link1Template = template.Must(template.ParseFiles("link1.html"))
var link2Template = template.Must(template.ParseFiles("link2.html"))

type userInfo struct {
	Name string `json:"Name"`
	Email string `json:"Email"`
	Message string `json:"Message"`
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/executeForm", formHandler)
	http.HandleFunc("/link1", link1Handler)
	http.HandleFunc("/link2", link2Handler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	expireDate := time.Now().AddDate(0, 0, 1)

	user := userInfo{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	userData, _ := json.Marshal(user)
	log.Printf("jSoN encode; %v", userData)
	for i, val := range userData {
		if val == 34 {
			userData[i] = 96
		}
	}

	cookie := http.Cookie {
		Name: "testCookie",
		Value: string(userData),
		Expires: expireDate,
		RawExpires: expireDate.Format(time.UnixDate),
		MaxAge: 86400,
		Secure: false,
		HttpOnly: false,
	}

	http.SetCookie(w, &cookie)

	err := formTemplate.Execute(w, struct{
				Name string
				Email string
				Message string
			}{r.FormValue("name"),r.FormValue("email"),r.FormValue("message")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func link1Handler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("name")+" "+r.FormValue("email")+" "+r.FormValue("message")
	err := link1Template.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func link2Handler(w http.ResponseWriter, r *http.Request) {
	var user userInfo
	cookie, _ := r.Cookie("testCookie")
	log.Printf("cookie2: %v", cookie.Value)

	byteSlice := []byte(cookie.Value)
	for i, val := range byteSlice{
		if val == 96 {
			byteSlice[i] = 34
		}
	}
	err := json.Unmarshal(byteSlice, &user)
	log.Print(user)
	err = link2Template.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
