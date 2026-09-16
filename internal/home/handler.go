package home

import (
	"net/http"
	"smply/app/render"
)

func Page(w http.ResponseWriter, r *http.Request) {
	render.Page(w, "home.html", render.ViewData{
		Title: "Home",
		Page:  "home",
	})
}

func PrivacyPage(w http.ResponseWriter, r *http.Request) {
	render.Page(w, "privacy.html", render.ViewData{
		Title: "Privacy Policy",
		Page:  "privacy",
	})
}

func TermsPage(w http.ResponseWriter, r *http.Request) {
	render.Page(w, "terms.html", render.ViewData{
		Title: "Terms of Service",
		Page:  "terms",
	})
}
