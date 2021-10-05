package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Group struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	Concerts     []ConcertData
	Lenght       int
}

type ConcertData struct {
	Location Place
	Date     []string
}

type Place struct {
	City    string
	Country string
}

type Data struct {
	Groups []Group
}

func main() {
	data := ParseJsonData()

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	http.HandleFunc("/", LoadMainPage(data))
	http.HandleFunc("/exit", ShutdownServer)

	log.Panic(http.ListenAndServe(":3030", nil))
}

func LoadMainPage(data Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mainTempl, err := template.ParseFiles("index.html")
		groupTempl, _ := template.ParseFiles("group-page.html")

		if err != nil {
			panic(err)
		}

		if id, err := strconv.Atoi(r.URL.Path[1:]); err == nil {
			if 0 < id && id <= len(data.Groups) {
				exactGroupData := data.Groups[id-1]
				exactGroupData.Lenght = len(data.Groups)

				if err := groupTempl.Execute(w, exactGroupData); err != nil {
					panic(err)
				}

				return
			}
		}

		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if err := mainTempl.Execute(w, data); err != nil {
			panic(err)
		}
	}
}

func ParseJsonData() Data {
	var result Data

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &result.Groups); err != nil {
		panic(err)
	}

	result = ParseConcerts(result)

	return result
}

func ParseConcerts(base Data) Data {
	for i := 0; i < len(base.Groups); i++ {
		var concertToAppend ConcertData
		var allConcerts []ConcertData
		var tempMap map[string]interface{}

		res, _ := http.Get(base.Groups[i].Relations)
		body, _ := ioutil.ReadAll(res.Body)

		if err := json.Unmarshal([]byte(body), &tempMap); err != nil {
			panic(err)
		}

		for i, v := range tempMap["datesLocations"].(map[string]interface{}) {
			tempDatesInOneString := strings.Trim(fmt.Sprint(v), "[]")
			tempDates := strings.Split(tempDatesInOneString, " ")
			tempSplited := strings.Split(i, "-")

			tempSplited[0] = strings.Title(strings.Replace(tempSplited[0], "_", " ", -1))
			tempSplited[1] = strings.ToUpper(strings.Replace(tempSplited[1], "_", " ", -1))

			concertToAppend.Location.City = tempSplited[0]
			concertToAppend.Location.Country = tempSplited[1]
			concertToAppend.Date = tempDates

			allConcerts = append(allConcerts, concertToAppend)
		}

		base.Groups[i].Concerts = allConcerts
	}

	return base
}

func ShutdownServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
