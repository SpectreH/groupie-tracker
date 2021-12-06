package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"groupie-tracker/models"
	"groupie-tracker/parser"
	"log"
	"net/http"
)

func main() {
	data := parser.ParseJsonData()
	models.InitModels()

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))
	js := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	repo := handlers.NewRepo(&data)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Handler)
	http.HandleFunc("/search", handlers.Repo.LoadSearch)
	http.HandleFunc("/exit", handlers.ShutdownServer)

	fmt.Println("Server is listening on port 3030...")
	log.Panic(http.ListenAndServe(":3030", nil))
}
