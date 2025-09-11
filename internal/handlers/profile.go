package handlers

import (
	"GoQuotes/internal/models"
	"GoQuotes/internal/templates"
	"GoQuotes/internal/utils"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ProfileHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	idv := session.Values["user_id"]

	uid, ok := utils.ToUintID(w, r, idv)

	if !ok {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	var user models.User
	if err := db.First(&user, uid).Error; err != nil {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	templates.Tmpl.ExecuteTemplate(w, "profile", map[string]string{
		"Username": user.Username,
	})

}
