package handlers

import (
	"GoQuotes/internal/models"
	"GoQuotes/internal/templates"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func RegisterHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.Tmpl.ExecuteTemplate(w, "register", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user models.User

		user.SetData(username, password)

		result := db.Create(&user)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Регистрация пользователя " + username + " прошла успешно")
	}
}

func LoginHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.Tmpl.ExecuteTemplate(w, "login", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user models.User

		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			templates.Tmpl.ExecuteTemplate(w, "login", map[string]string{
				"Error": "Неверное имя пользователя",
			})
			return

		}

		if user.CheckPassword(password) {
			session, _ := store.Get(r, "session")
			session.Values["user_id"] = int(user.ID)
			if err := session.Save(r, w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			templates.Tmpl.ExecuteTemplate(w, "login", map[string]string{
				"Error": "Неверный пароль",
			})
			return
		}
		fmt.Println("Пользователь " + username + " успешно зашел!")

	}
}
