package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"url-shortener-go/models"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func ShortenURL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.URL
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &input)

		code := generateCode(6)
		_, err := db.Exec("INSERT INTO urls (code, original_url) VALUES (?, ?)", code, input.OriginalURL)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		response := models.URL{Code: code, OriginalURL: input.OriginalURL}
		json.NewEncoder(w).Encode(response)
	}
}

func RedirectURL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[1:]
		var url string
		err := db.QueryRow("SELECT original_url FROM urls WHERE code = ?", code).Scan(&url)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
}
