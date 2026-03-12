package main

import (
	"log"
	"net/http"
	"snipfyi/config"
	"snipfyi/handlers"
)

func main() {

	envErr := config.LoadEnv()
	if envErr != nil {
		log.Fatal(envErr)
	}

	dbErr := config.InitDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/shorten", handlers.ShortenPage)

	http.HandleFunc("POST /api/shorten", handlers.Shorten)
	http.HandleFunc("/r/{code}", handlers.Redirect)
	http.HandleFunc("/r/{code}/stats", handlers.Stats)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
