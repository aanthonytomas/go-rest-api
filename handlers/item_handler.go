package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"go-rest-api/models"
)

// GetItems returns a handler function that fetches all items from the database
func GetItems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Fetch items from database
		items, err := models.FetchItems(db)
		if err != nil {
			log.Printf("Error fetching items: %v", err)
			http.Error(w, `{"error": "Failed to fetch items"}`, http.StatusInternalServerError)
			return
		}

		// Handle case when no items are found
		if len(items) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Encode and send the response
		if err := json.NewEncoder(w).Encode(items); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		}
	}
}

// AddItem handles POST requests to create a new item
func AddItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Parse request body
		var item models.Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		// Insert into database (you'll need to implement this function in models)
		// if err := models.CreateItem(db, &item); err != nil {
		// 	log.Printf("Error creating item: %v", err)
		// 	http.Error(w, `{"error": "Failed to create item"}`, http.StatusInternalServerError)
		// 	return
		// }


		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	}
}
