package utils

import (
	"GoQuotes/internal/models"
	"GoQuotes/internal/templates"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ToUintID(idv interface{}) (uint, bool) {
	var uid uint

	switch v := idv.(type) {
	case int:
		uid = uint(v)
	case int64:
		uid = uint(v)
	case uint:
		uid = v
	case uint64:
		uid = uint(v)
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			uid = uint(n)
		} else {
			return 0, false
		}
	default:
		return 0, false
	}

	return uid, true
}

func Login(w http.ResponseWriter, r *http.Request, res *gorm.DB, user models.User, store *sessions.CookieStore, username string, password string) {
	if res.Error != nil {
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

func IsLoggedIn(store *sessions.CookieStore, r *http.Request) (uint, bool) {
	session, _ := store.Get(r, "session")
	uid, ok := ToUintID(session.Values["user_id"])

	if !ok {

		return 0, false
	} else {
		return uid, true
	}

}
