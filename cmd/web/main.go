package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gnn/internal/app"
)

func main() {
	addr := envOrDefault("ADDR", ":8080")
	logger := log.New(os.Stdout, "", log.LstdFlags)
	server := app.NewServer(logger)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           server.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	logger.Printf("gnn listening on http://localhost%s", addr)

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
