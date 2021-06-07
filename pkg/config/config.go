package config

import "html/template"

// AppConfig hold the configuration of the app
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
