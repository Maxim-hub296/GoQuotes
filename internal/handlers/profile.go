package handlers

import (
	"GoQuotes/internal/models"
	"GoQuotes/internal/templates"
	"GoQuotes/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ProfileHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {

	uid, ok := utils.IsLoggedIn(store, r)

	if !ok {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	var user models.User
	if err := db.First(&user, uid).Error; err != nil {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	var quotesCount int64
	db.Model(&models.Quote{}).Where("user_id = ?", user.ID).Count(&quotesCount)

	templates.Tmpl.ExecuteTemplate(w, "profile", map[string]string{
		"Username":  user.Username,
		"QuotCount": strconv.FormatInt(quotesCount, 10),
	})

}
