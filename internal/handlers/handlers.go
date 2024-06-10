package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/erichRoberts/bookings/internal/config"
	"github.com/erichRoberts/bookings/internal/models"
	"github.com/erichRoberts/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About presents the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "this is from the test"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Bigroom presents the Bigroom page
func (m *Repository) Bigroom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "bigroom.page.tmpl", &models.TemplateData{})
}

// Littleroom present the Littleroom page
func (m *Repository) Littleroom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "littleroom.page.tmpl", &models.TemplateData{})
}

// Availability presents the search-availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles the search-availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start is %s, end is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println((err))
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write((out))
}

// Contact present the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
