package spread

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetSymbolsHandler(w http.ResponseWriter, r *http.Request) {

	// Only GET method allowed for this endpoint
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Symbols)

	// Check for encoding response
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func GetSpreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := strings.TrimPrefix(r.URL.Path, "/spreads/")
	if symbol == "" {
		http.NotFound(w, r)
		return
	}

	spread, found := GetSpread(symbol)
	if !found {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(spread)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}