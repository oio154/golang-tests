package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Uruchamiam obliczenia...")
	
	// Prosta operacja matematyczna w pętli
	var wynik float64
	for i := 1.0; i <= 1000.0; i++ {
		wynik += math.Sqrt(i)
	}

	fmt.Printf("Wynik końcowy obliczeń: %.2f\n", wynik)
	fmt.Println("Aplikacja kończy działanie. Czekam na zamknięcie...")
	
	// Utrzymujemy proces przy życiu, żeby kontener się od razu nie wyłączył
	for {
		time.Sleep(1 * time.Hour)
	}
}
