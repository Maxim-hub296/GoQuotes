package routes

import (
	"GoQuotes/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func QuoteRoutes(r *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateQuote(db, store, w, r)
	})
}
