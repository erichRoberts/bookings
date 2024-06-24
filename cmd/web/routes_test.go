package main

import (
	"fmt"
	"testing"

	"github.com/erichRoberts/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing; test passed
	default:
		t.Errorf(fmt.Sprintf("Type is not *chi.mux. Type is %T", v))

	}

}
