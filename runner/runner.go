package main

import (
	"fmt"
	"net/http"

	"../../elizago"
)

var counter = 0

//get called when user submit request by ajax call
func userinputhandler(w http.ResponseWriter, r *http.Request) {

	//get user input
	input := r.URL.Query().Get("value")

	//first time called,give some greetings
	if counter < 1 {

		fmt.Fprintf(w, "Eliza: %s", elizago.ElizaHi())

		counter++

	} else {

		//write response back
		fmt.Fprintf(w, "Eliza: %s", elizago.ReplyTo(input))
	}

}

//The main entry
func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	//serve static html file
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	//handle request /user-input and ajax get method
	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)

}
