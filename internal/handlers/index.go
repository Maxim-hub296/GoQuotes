package handlers

import (
	"GoQuotes/internal/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Tmpl.ExecuteTemplate(w, "privacy", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
