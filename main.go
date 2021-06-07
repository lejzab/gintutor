package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on address %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
