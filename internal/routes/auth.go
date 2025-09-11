package routes

import (
	"GoQuotes/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func AuthRoutes(r *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(db, store, w, r)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(db, store, w, r)
	})

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(store, w, r)
	})

	r.HandleFunc("/delete_user", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserHandler(db, store, w, r)
	})
}
