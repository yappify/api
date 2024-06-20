package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yappify/api/internal/database"
)

func main() {
	port, dbURL := loadEnv()

	if port == "" {
		port = "8000"
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	api := Config{
		DB: database.New(conn),
	}

	// Define http server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: api.routes(),
	}

	log.Printf("Starting server on port %s...\n", port)

	// Start http server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func loadEnv() (string, string) {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	return port, dbURL
}
