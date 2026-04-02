package spread

import (
	"encoding/json"
	"net/http"
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