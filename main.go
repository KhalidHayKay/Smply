package main

import (
	"log"
	"net/http"
	"snipfyi/config"
	"snipfyi/handlers"
)

func main() {

	config.LoadEnv()

	err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Page routes
	http.HandleFunc("GET /", handlers.Home)
	http.HandleFunc("GET /shorten", handlers.ShortenPage)
	http.HandleFunc("GET /stats/{code}", handlers.Stats)
	http.HandleFunc("GET /{code}", handlers.Redirect)

	// API routes
	http.HandleFunc("POST /api/shorten", handlers.Shorten)
	http.HandleFunc("GET /api/stats/{code}", handlers.StatsApi)
	http.HandleFunc("GET /api/{code}", handlers.RedirectApi)

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	port := config.Env.AppPort

	if port == "" && config.Env.AppEnv == "development" {
		port = "8000"
	}

	log.Printf("Starting server on port %s in %s mode", port, config.Env.AppEnv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
