package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/erichRoberts/bookings/internal/config"
	"github.com/erichRoberts/bookings/internal/handlers"
	"github.com/erichRoberts/bookings/internal/render"
)

// portNumber is the port that will be listened to
const portNumber = ":8080"

// app is the config for the current session
var app config.AppConfig

// session is a pointer to the scs.SessionManager
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	println("starting")

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create the themplate cache and store it in the app
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	println("created template cache")
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// start the server
	println("Server listening on port ", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
