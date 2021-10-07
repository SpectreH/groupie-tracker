package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

type ConcertData struct {
	Location Place
	Date     []string
}

type Place struct {
	City    string
	Country string
}

type Data struct {
	Groups                  []Group
	ConcertPlaces           []string
	InStandardSintLocations []string
}

type Date struct {
	Day   int
	Month int
	Year  int
}

func main() {
	data := ParseJsonData()

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	http.HandleFunc("/", LoadPage(data))
	http.HandleFunc("/search", LoadPage(data))
	http.HandleFunc("/exit", ShutdownServer)

	log.Panic(http.ListenAndServe(":3030", nil))
}

func LoadPage(data Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mainTempl, err := template.ParseFiles("index.html")
		groupTempl, _ := template.ParseFiles("group-page.html")
		searchTempl, _ := template.ParseFiles("search.html")

		if err != nil {
			panic(err)
		}

		if r.URL.Path == "/search" {
			if err := searchTempl.Execute(w, data); err != nil {
				panic(err)
			}

			return
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
	result = ParseDates(result)

	return result
}

func ParseConcerts(base Data) Data {
	var concertPlaces []string
	var standSintlocations []string

	for i := 0; i < len(base.Groups); i++ {
		var concertToAppend ConcertData
		var standSintConcerts []string
		var allConcerts []ConcertData
		var tempMap map[string]interface{}

		res, _ := http.Get(base.Groups[i].Relations)
		body, _ := ioutil.ReadAll(res.Body)

		if err := json.Unmarshal([]byte(body), &tempMap); err != nil {
			panic(err)
		}

		for i, v := range tempMap["datesLocations"].(map[string]interface{}) {
			var concertPlaceToAppend string

			standSintConcerts = append(standSintConcerts, i)

			tempDatesInOneString := strings.Trim(fmt.Sprint(v), "[]")
			tempDates := strings.Split(tempDatesInOneString, " ")
			tempLocationSplited := strings.Split(i, "-")

			tempLocationSplited[0] = strings.Title(strings.Replace(tempLocationSplited[0], "_", " ", -1))
			tempLocationSplited[1] = strings.ToUpper(strings.Replace(tempLocationSplited[1], "_", " ", -1))

			concertToAppend.Location.City = tempLocationSplited[0]
			concertToAppend.Location.Country = tempLocationSplited[1]
			concertToAppend.Date = tempDates

			concertPlaceToAppend = tempLocationSplited[0] + ", " + tempLocationSplited[1]

			allConcerts = append(allConcerts, concertToAppend)
			concertPlaces = append(concertPlaces, concertPlaceToAppend)
			standSintlocations = append(standSintlocations, i)
		}

		concertPlaces = RemoveDuplicateStr(concertPlaces)
		standSintlocations = RemoveDuplicateStr(standSintlocations)
		standSintConcerts = RemoveDuplicateStr(standSintConcerts)

		base.Groups[i].InStandardSintConcertLocations = standSintConcerts
		base.Groups[i].Concerts = allConcerts
	}

	sort.Strings(concertPlaces)
	base.ConcertPlaces = concertPlaces
	base.InStandardSintLocations = standSintlocations

	return base
}

func ParseDates(base Data) Data {
	for i := 0; i < len(base.Groups); i++ {
		var dateToAppend Date

		tempDate := strings.Split(base.Groups[i].FirstAlbum, "-")

		dateToAppend.Day, _ = strconv.Atoi(tempDate[0])
		dateToAppend.Month, _ = strconv.Atoi(tempDate[1])
		dateToAppend.Year, _ = strconv.Atoi(tempDate[2])

		base.Groups[i].FirstAlbumSep = dateToAppend
	}

	return base
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ShutdownServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
