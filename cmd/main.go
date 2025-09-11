package main

import (
	"GoQuotes/internal/database"
	"GoQuotes/internal/routes"
	"GoQuotes/internal/templates"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {

	fmt.Println("Подготовка бд...")
	db := database.InitDB()
	fmt.Println("Подготовка хранилища сессий")
	store := sessions.NewCookieStore([]byte("super-secret-password"))
	store.MaxAge(600 * 30)
	fmt.Println("Загружаем шаблоны...")
	templates.LoadTemplates()

	mainRouter := mux.NewRouter()

	routes.IndexRoutes(mainRouter)

	authSubRouter := mainRouter.PathPrefix("/auth").Subrouter()
	routes.AuthRoutes(authSubRouter, db, store)

	fmt.Println("Сервер работает на http://localhost:8181")

	err := http.ListenAndServe(":8181", mainRouter)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ")
		fmt.Println(err)
	}

}
