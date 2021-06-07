package main

import (
	"fmt"
	"github.com/lejzab/gintutor/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on address %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
