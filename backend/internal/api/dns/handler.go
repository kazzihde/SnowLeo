package dnsapi

import (
	"encoding/json"
	"net/http"

	"github.com/datashelll/SnowLeo/internal/services/dns"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")

	if domain == "" {
		http.Error(w, "missing domain parameter", http.StatusBadRequest)
		return
	}

	result := dns.Lookup(domain)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
