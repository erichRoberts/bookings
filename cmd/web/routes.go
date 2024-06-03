package main

import (
	"net/http"

	"github.com/erichRoberts/bookings/pkg/config"
	"github.com/erichRoberts/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/bigroom", handlers.Repo.Bigroom)
	mux.Get("/littleroom", handlers.Repo.Littleroom)
	mux.Get("/booknow", handlers.Repo.Booknow)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
