package main

import (
	"log"
	"net/http"
	"snipfyi/config"
	"snipfyi/handlers"
)

func main() {

	err := config.LoadEnv()
	if err != nil {
		log.Println(err)
	}

	err = config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/shorten", handlers.ShortenPage)

	http.HandleFunc("POST /api/shorten", handlers.Shorten)
	http.HandleFunc("/r/{code}", handlers.Redirect)
	http.HandleFunc("/r/{code}/stats", handlers.Stats)

	port := config.Env.AppPort

	if port == "" && config.Env.AppEnv == "development" {
		port = "8000"
	}

	log.Printf("Starting server on port %s in %s mode", port, config.Env.AppEnv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
