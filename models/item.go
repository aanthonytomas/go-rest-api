package models

import (
	"database/sql"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FetchItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
