package whoisapi

import (
	"encoding/json"
	"net/http"

	"github.com/datashelll/SnowLeo/internal/services/whois"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}

	result, err := whois.Lookup(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
