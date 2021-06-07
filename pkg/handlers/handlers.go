package handlers

import (
	"github.com/lejzab/gintutor/pkg/config"
	"github.com/lejzab/gintutor/pkg/render"
	"net/http"
)

//Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository for handlers
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		appConfig,
	}
}

// NewHandlers set new repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
