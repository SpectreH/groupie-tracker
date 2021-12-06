package models

import (
	"groupie-tracker/render"
	"text/template"
)

// TemplateData holds data for handler
type TemplateData struct {
	Groups                  []Group
	ConcertPlaces           []string
	InStandardSintLocations []string
}

// Group holds one group data
type Group struct {
	Id                             int      `json:"id"`
	Image                          string   `json:"image"`
	Name                           string   `json:"name"`
	Members                        []string `json:"members"`
	CreationDate                   int      `json:"creationDate"`
	FirstAlbum                     string   `json:"firstAlbum"`
	Relations                      string   `json:"relations"`
	FirstAlbumSep                  Date
	Concerts                       []ConcertData
	InStandardSintConcertLocations []string
	Lenght                         int
}

// ConcertData holds concert data of one group
type ConcertData struct {
	Location Place
	Date     []string
}

// Place holds concert location data
type Place struct {
	City    string
	Country string
}

// Date holds concert date
type Date struct {
	Day   int
	Month int
	Year  int
}

// TEMPLATES holds all templates in a map
var TEMPLATES map[string]*template.Template

func InitModels() {
	TEMPLATES = render.ParseTemplates()
}
