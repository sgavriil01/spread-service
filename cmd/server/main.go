package main

import (
	"log"
	"net/http"
	"github.com/sgavriil01/spread-service/internal/spread"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/symbols", spread.GetSymbolsHandler)
	mux.HandleFunc("/spreads/", spread.SpreadHandler)

	log.Println("server running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}