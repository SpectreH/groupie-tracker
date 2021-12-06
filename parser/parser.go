package parser

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// ParseJsonData parses json from the link and returns as TemplateData
func ParseJsonData() models.TemplateData {
	var result models.TemplateData

	res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &result.Groups); err != nil {
		panic(err)
	}

	result = ParseConcerts(result)
	result = ParseDates(result)

	return result
}

// ParseConcerts parses all conserts from json and returns them
func ParseConcerts(base models.TemplateData) models.TemplateData {
	var concertPlaces []string
	var standSintlocations []string

	for i := 0; i < len(base.Groups); i++ {
		var concertToAppend models.ConcertData
		var standSintConcerts []string
		var allConcerts []models.ConcertData
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

// ParseDates parses all conserts dates from json and returns them
func ParseDates(base models.TemplateData) models.TemplateData {
	for i := 0; i < len(base.Groups); i++ {
		var dateToAppend models.Date

		tempDate := strings.Split(base.Groups[i].FirstAlbum, "-")

		dateToAppend.Day, _ = strconv.Atoi(tempDate[0])
		dateToAppend.Month, _ = strconv.Atoi(tempDate[1])
		dateToAppend.Year, _ = strconv.Atoi(tempDate[2])

		base.Groups[i].FirstAlbumSep = dateToAppend
	}

	return base
}

// RemoveDuplicateStr removes dublicates
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
