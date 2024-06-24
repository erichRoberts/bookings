package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds application config data
type AppConfig struct {
	// UseCache determines if the template cache will be used or if new requestes will be rendered afresh
	UseCache bool

	// TempateCache is a cache of templates that have already been parsed and are ready to be executed
	TemplateCache map[string]*template.Template

	// InfoLog is the Information level logger
	InfoLog *log.Logger

	// ErrorLog is the Error level logger
	ErrorLog *log.Logger

	// InProduction should be set to true for production
	InProduction bool

	Session *scs.SessionManager
}
