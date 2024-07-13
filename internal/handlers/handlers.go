package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/motoshi-suzuki-hp/bookings/internal/config"
	"github.com/motoshi-suzuki-hp/bookings/internal/forms"
	"github.com/motoshi-suzuki-hp/bookings/internal/models"
	"github.com/motoshi-suzuki-hp/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Reposiory

// Repository is the repository type
type Reposiory struct {
	App *config.AppConfig
}

// NewRepo coreates a new repository
func NewRepo(a *config.AppConfig) *Reposiory {
	return &Reposiory{
		App: a,
	}
}

// NewHAndlers sets the repository for the handlers
func NewHandlers(r *Reposiory) {
	Repo = r
}


// Home is the home page handler
func (m *Reposiory) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Reposiory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	
	
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays from
func (m *Reposiory) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation renders the posting of the reservation from
func (m *Reposiory) PostReservation(w http.ResponseWriter, r *http.Request) {

}

// Generals renders the room page
func (m *Reposiory) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Reposiory) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the availability page
func (m *Reposiory) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the availability page
func (m *Reposiory) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")


	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}


type jsonResponse struct {
	OK bool `json:"ok`
	Message string `json:"message`
}
// AvailabilityJSON handles request fo availability and send JSON response
func (m *Reposiory) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	// log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the availability page
func (m *Reposiory) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}