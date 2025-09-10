package main

import (
	"GoQuotes/internal/database"
	"GoQuotes/internal/routes"
	"GoQuotes/internal/templates"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Подготовка бд...")
	database.InitDB()

	fmt.Println("Загужаем шаблоны...")
	templates.LoadTemplates()

	mainRouter := mux.NewRouter()

	routes.IndexRoutes(mainRouter)

	fmt.Println("Сервер работает на http://localhost:8181")

	err := http.ListenAndServe(":8181", mainRouter)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ")
		fmt.Println(err)
	}

}
