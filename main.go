package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-rest-api/cache"
	"go-rest-api/db"
	"go-rest-api/handlers"
)

func main() {
	// Initialize database and cache
	dbConn := db.InitPostgres()
	cache.InitRedis()

	// Setup router
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/items", handlers.GetItems(dbConn)).Methods("GET")
	api.HandleFunc("/items", handlers.AddItem(dbConn)).Methods("POST")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	// Start server
	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
