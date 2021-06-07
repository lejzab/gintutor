package main

import (
	"fmt"
	"github.com/lejzab/gintutor/pkg/config"
	"github.com/lejzab/gintutor/pkg/handlers"
	"github.com/lejzab/gintutor/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cpuld not create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on address %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
