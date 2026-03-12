package views

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, page string, data any) {
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles(
		"templates/layouts/layout.html",
		"templates/pages/"+page,
	)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		w.Write([]byte(err.Error()))
	}
}
