package handlers

import (
	"GoQuotes/internal/models"
	"GoQuotes/internal/templates"
	"GoQuotes/internal/utils"
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
			templates.Tmpl.ExecuteTemplate(w, "register", map[string]string{
				"Error": "Имя пользователя уже занято",
			})
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
		http.Redirect(w, r, "/profile", http.StatusSeeOther)

	}
}

func LogoutHandler(store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteUserHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	idv := session.Values["user_id"]
	uid, _ := utils.ToUintID(w, r, idv)

	var user models.User
	if err := db.First(&user, uid).Error; err != nil {
		fmt.Println("Пользователь не найден:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := user.Delete(db); err != nil {
		fmt.Println("Ошибка при удалении:", err)
	} else {
		fmt.Println("Пользователь удалён")
	}

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
