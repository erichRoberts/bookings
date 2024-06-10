package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/erichRoberts/bookings/pkg/config"
	"github.com/erichRoberts/bookings/pkg/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// NewTemplates set the config for the new template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the TemplateCache from the config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	AddDefaultData(td, r)

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)

	// render the template

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates a cache of all parsed templates
func CreateTemplateCache() (map[string]*template.Template, error) {

	templateDir := "./templates/"
	pageNames := "*.page.tmpl"
	layoutNames := "*.layout.tmpl"
	myCache := make(map[string]*template.Template)

	// get all of the files called *.page.tmpl
	pages, err := filepath.Glob(templateDir + pageNames)
	if err != nil {
		return myCache, err
	}

	// get all the layout files
	layouts, err := filepath.Glob(templateDir + layoutNames)
	if err != nil {
		return myCache, err
	}

	// range through all files ending page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(templateDir + layoutNames)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

// This is the pre load for templates

// // RenderTemplate renders templates using html/template
// func RenderTemplate(w http.ResponseWriter, tmpl string) {

// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error writing template ", tmpl, err)
// 	}
// }

// This is lazy load for templates
// // tc is the template cache
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// we need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = creatTemplate(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func creatTemplate(t string) error {
// 	templates := []string{
// 		"./templates/" + t,
// 		// he has sprintf("./templates/%s",s)
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache
// 	tc[t] = tmpl
// 	return nil
// }
