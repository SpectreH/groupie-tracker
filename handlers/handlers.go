package handlers

import (
	"groupie-tracker/models"
	"net/http"
	"os"
	"strconv"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	GroupsData *models.TemplateData
}

// NewRepo creates a new repository
func NewRepo(a *models.TemplateData) *Repository {
	return &Repository{
		GroupsData: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Handler redirects between main and group pages
func (d Repository) Handler(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.URL.Path[1:]); err == nil {
		if 0 < id && id <= len(d.GroupsData.Groups) {
			exactGroupData := d.GroupsData.Groups[id-1]
			exactGroupData.Lenght = len(d.GroupsData.Groups)

			if err := models.TEMPLATES["group-page.html"].Execute(w, exactGroupData); err != nil {
				panic(err)
			}

			return
		}
	}

	if r.URL.Path != "/" {
		d.LoadError(w, r)
		return
	}

	d.LoadMain(w, r)
}

// LoadMain loads main page
func (d Repository) LoadMain(w http.ResponseWriter, r *http.Request) {
	if err := models.TEMPLATES["index.html"].Execute(w, d); err != nil {
		panic(err)
	}
}

// LoadSearch loads search page
func (d Repository) LoadSearch(w http.ResponseWriter, r *http.Request) {
	if err := models.TEMPLATES["search.html"].Execute(w, d); err != nil {
		panic(err)
	}
}

// LoadError loads 404 error page
func (d Repository) LoadError(w http.ResponseWriter, r *http.Request) {
	if err := models.TEMPLATES["error.html"].Execute(w, d); err != nil {
		panic(err)
	}
}

// ShutdownServer shut down server
func ShutdownServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
