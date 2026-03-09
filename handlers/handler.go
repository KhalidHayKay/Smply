package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"shortener/service"
	"shortener/utils"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("url")

	if input == "" {
		Error(w, http.StatusUnprocessableEntity, "'url' is a required field")
		return
	}

	if !utils.IsValidURL(input) {
		Error(w, http.StatusUnprocessableEntity, "Url not valid")
		return
	}

	result, err := service.StoreUrl(input)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	result.ShortToUrl()

	Success(w, http.StatusCreated, result)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("redirect")

	url, err := service.Retrieve(short)

	if err != nil {
		if err == sql.ErrNoRows {
			Error(w, http.StatusNotFound, "Not found")
			return
		}

		Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	go func() {
		err := service.IncrementVisited(url.Id)

		if err != nil {
			log.Println(err)
		}
	}()

	http.Redirect(w, r, url.Original, http.StatusFound)
}

func Stats(w http.ResponseWriter, r *http.Request) {
	//
}
