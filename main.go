package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Foo struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Data struct {
	Foos []Foo
}

func main() {
	data := ParseJsonData()

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", LoadMainPage(data))
	http.HandleFunc("/exit", ShutdownServer)

	log.Panic(http.ListenAndServe(":3030", nil))
}

func LoadMainPage(data Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("index.html")

		if err != nil {
			panic(err)
		}

		if err := templ.Execute(w, data); err != nil {
			panic(err)
		}

	}
}

func ParseJsonData() Data {
	var result Data

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &result.Foos); err != nil {
		panic(err)
	}

	return result
}

func ShutdownServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
