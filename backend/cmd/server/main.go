package main

import (
	"fmt"
	"log"
	"net/http"

	dnsapi "github.com/datashelll/SnowLeo/internal/api/dns"
	ipapi "github.com/datashelll/SnowLeo/internal/api/ip"
	whoisapi "github.com/datashelll/SnowLeo/internal/api/whois"
	"github.com/datashelll/SnowLeo/internal/config"
	"github.com/datashelll/SnowLeo/internal/database"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	if cfg.DBUser == "" || cfg.DBPassword == "" {
		log.Fatal("missing DB credentials in environment")
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/dns", dnsapi.Handler)
	http.HandleFunc("/api/v1/ip", ipapi.Handler)
	http.HandleFunc("/api/v1/whois", whoisapi.Handler)

	fmt.Println("SnowLeo backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
