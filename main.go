package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var jsonData = `[{"id":1,"image":"https://groupietrackers.herokuapp.com/api/images/queen.jpeg","name":"Queen","members":["Freddie Mercury","Brian May","John Daecon","Roger Meddows-Taylor","Mike Grose","Barry Mitchell","Doug Fogie"],"creationDate":1970,"firstAlbum":"14-12-1973","locations":"https://groupietrackers.herokuapp.com/api/locations/1","concertDates":"https://groupietrackers.herokuapp.com/api/dates/1","relations":"https://groupietrackers.herokuapp.com/api/relation/1"},
{"id":2,"image":"https://groupietrackers.herokuapp.com/api/images/soja.jpeg","name":"SOJA","members":["Jacob Hemphill","Bob Jefferson","Ryan \"Byrd\" Berty","Ken Brownell","Patrick O'Shea","Hellman Escorcia","Rafael Rodriguez","Trevor Young"],"creationDate":1997,"firstAlbum":"05-06-2002","locations":"https://groupietrackers.herokuapp.com/api/locations/2","concertDates":"https://groupietrackers.herokuapp.com/api/dates/2","relations":"https://groupietrackers.herokuapp.com/api/relation/2"}]`

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
	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	body, _ := ioutil.ReadAll(res.Body)

	var d Data
	if err := json.Unmarshal([]byte(body), &d.Foos); err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("index.html")

		if err != nil {
			panic(err)
		}

		if err := templ.Execute(rw, d); err != nil {
			panic(err)
		}
	})
	log.Panic(http.ListenAndServe(":3030", nil))
}
