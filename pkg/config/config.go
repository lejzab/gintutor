package config

import "html/template"

// AppConfig hold the configuration of the app
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
