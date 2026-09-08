package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message     string `json:"message"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Time        string `json:"time"`
}

func main() {
	port := getEnv("PORT", "8080")
	environment := getEnv("ENVIRONMENT", "dev")
	version := getEnv("APP_VERSION", "local")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message:     "Hello from Go on AKS!",
			Environment: environment,
			Version:     version,
			Time:        time.Now().UTC().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("READY"))
	})

	log.Printf("Starting server on port %s", port)
	log.Printf("Environment: %s", environment)
	log.Printf("Version: %s", version)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
