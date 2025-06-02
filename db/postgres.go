package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() *sql.DB {
	var err error
	// First try connecting with the application user
	connStr := "host=localhost port=5432 user=sampleuser password=samplepass dbname=sampledb sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Postgres connection failed:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Postgres ping failed:", err)
	}
	log.Println("Connected to Postgres!")
	return DB
}
