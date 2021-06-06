package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {

}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Staring application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
