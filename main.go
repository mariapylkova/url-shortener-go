package main

import (
	"log"
	"net/http"
	"url-shortener-go/database"
	"url-shortener-go/handlers"
)

func main() {
	db := database.InitDB("urlshortener.db")
	defer db.Close()

	http.HandleFunc("/shorten", handlers.ShortenURL(db))
	http.HandleFunc("/", handlers.RedirectURL(db))

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
