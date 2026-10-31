package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Definiujemy zestaw znaków alfanumerycznych
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Funkcja generująca losowy ciąg znaków o określonej długości
func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Ustawiamy nagłówek na czysty tekst zamiast HTML
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Generujemy 10 000 losowych znaków
	randomData := generateRandomString(10000)
	
	// Wypisujemy wynik do odpowiedzi HTTP
	fmt.Fprint(w, randomData)
}

func main() {
	// Inicjalizujemy generator liczb losowych aktualnym czasem
	rand.Seed(time.Now().UnixNano())

	// Obsługujemy każdy adres URI (pusta ścieżka "/" łapie wszystko)
	http.HandleFunc("/", handleRequest)

	fmt.Println("Serwer WWW uruchomiony na porcie 8081...")
	
	// Uruchamiamy serwer na porcie 8081. 
	// Ta funkcja blokuje działanie programu i utrzymuje kontener przy życiu.
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("Błąd serwera: %v\n", err)
	}
}
