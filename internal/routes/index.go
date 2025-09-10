package routes

import (
	"GoQuotes/internal/handlers"

	"github.com/gorilla/mux"
)

func IndexRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/privacy", handlers.PrivacyHandler)
}
