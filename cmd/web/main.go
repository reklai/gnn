package main

import (
	"log"
	"net/http"
	"os"

	"gnn/internal/app"
)

func main() {
	addr := envOrDefault("ADDR", ":8080")
	logger := log.New(os.Stdout, "", log.LstdFlags)
	server := app.NewServer(logger)

	logger.Printf("gnn listening on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		logger.Fatal(err)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
