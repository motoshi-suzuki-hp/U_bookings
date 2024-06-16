package handlers

import (
	"net/http"

	"github.com/motoshi-suzuki-hp/bookings/pkg/config"
	"github.com/motoshi-suzuki-hp/bookings/pkg/models"
	"github.com/motoshi-suzuki-hp/bookings/pkg/render"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Reposiory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	
	
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


