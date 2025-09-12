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

type Quotes struct {
	Quotes []models.Quote
}

func CreateQuoteHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {

	uid, ok := utils.IsLoggedIn(store, r)

	if !ok {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}

	if r.Method == "GET" {
		err := templates.Tmpl.ExecuteTemplate(w, "quote_form", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		author := r.FormValue("author")
		text := r.FormValue("text")

		var quote models.Quote

		quote.Create(author, text, uid)

		result := db.Create(&quote)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		templates.Tmpl.ExecuteTemplate(w, "quote_form", map[string]string{
			"Success": "Цитата успешно сохранена",
		})
		fmt.Println("Цитата успешно сохранена!")

	}
}

func UserQuotesHandler(db *gorm.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	uid, ok := utils.IsLoggedIn(store, r)
	if !ok {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	var quotes []models.Quote
	db.Where("user_id = ?", uid).Find(&quotes)

	templates.Tmpl.ExecuteTemplate(w, "quotes", Quotes{
		Quotes: quotes,
	})

}
