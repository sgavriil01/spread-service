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

func PatchSpreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := strings.TrimPrefix(r.URL.Path, "/spreads/")
	if symbol == "" {
		http.NotFound(w, r)
		return
	}

	var req SetSpreadRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Spread <= 0 {
		http.Error(w, "spread must be greater than 0", http.StatusBadRequest)
		return
	}

	savedSpread, ok := SetSpread(symbol, req.Spread)
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(savedSpread)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func SpreadHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetSpreadHandler(w, r)

	case http.MethodPatch:
		PatchSpreadHandler(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}