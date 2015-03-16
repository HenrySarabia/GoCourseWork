//Author: Henry Sarabia
//Date: 3/15/2015
//This is a simple program that will create a letter based on a template and each
//specific client's information such as their name, honorific, attendance, donation
//and other upcoming events.

package main

import "text/template"
import "os"
import "log"

//This is the struct we will be using that will hold each client's information
//It's made up of 2 strings, 2 bools, and a slice of strings
type Client struct{
	Honorific string
	SurName string
	DidAttend bool
	DidDonate bool
	Events []string
}

func main() {

	//This is the constant that we will use as a template for the outgoing letter
	//There are several conditional statements to personalize each letter depending on the client's information
	const letter = `
	Dear {{.Honorific}} {{.SurName}},
	{{if .DidAttend}}
	Thank you so much for participating in the fundraiser.
	It was lovely to have you there and we sincerely hope you had a wonderful time.
	{{if.DidDonate}} We appreciate your generous contribution and guarantee that it will be used responsibly. {{end}}
	{{else}}
	We are so sorry you couldn't attend the fundraiser.
	The event was quite wonderful and we hope you can participate in the future.
	{{if .DidDonate}} Although you couldn't make it, we do appreciate your generous donation through proxy. We guarantee this will go to a great cause. {{end}}
	{{end}}
	We hope you can make it to these upcoming events:
	{{range .Events}}
	{{.}}{{end}}

	Warm regards,
	The Aperture Science Foundation
	`
	//This is a slice a strings that holds the upcoming events data
	var EventList = []string{"Bora Bora Boar Hunt", "Pasta La Vista", "The Million Mime March"}

	//This is a slice of Client types that will hold each client's information
	var clientList = []Client {
		{"Dr.", "Acula", true, true, EventList},
		{"Mr.", "Mister", false, true, EventList},
		{"Mrs.", "Doubtfire", true, false, EventList},
		{"Ms.", "Take",false, false, EventList},
	}

	//This is the template we create to parse the letter into
	t := template.Must(template.New("letter").Parse(letter))

	//This is the for loop that will iterate through each of the clients and form each of their letters
	//The function will output a log of an error if one happens to occur.
	for _, r := range clientList {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

// FROM: https://github.com/GoesToEleven/Web_Programming_CSUF/blob/master/src/github.com/goestoeleven/bwpwg-master/10_wedding_form_letter.go
