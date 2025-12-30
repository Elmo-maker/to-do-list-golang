package config

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// LOAD .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := "host=localhost port=5432 user=postgres password=270605 dbname=todo_db sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon:", err)
	}

	fmt.Println("Berhasil terhubung ke database!")
}
