package routes

import (
	"GoQuotes/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ProfileRoutes(r *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		handlers.ProfileHandler(db, store, w, r)
	})
}
