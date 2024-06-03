package handlers

import (
	"net/http"

	"github.com/erichRoberts/bookings/pkg/config"
	"github.com/erichRoberts/bookings/pkg/models"
	"github.com/erichRoberts/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	app *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// NewHandlers sets the Repository for the new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home presents the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About presents the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "this is from the test"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Bigroom presents the Bigroom page
func (m *Repository) Bigroom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "bigroom.page.tmpl", &models.TemplateData{})
}

// Littleroom present the Littleroom page
func (m *Repository) Littleroom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "littleroom.page.tmpl", &models.TemplateData{})
}

// Booknow presents the search-availability page
func (m *Repository) Booknow(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact present the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
