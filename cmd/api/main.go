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

const DEFAULT_PORT = "8000"
const DEVELOPMENT_DB_URL = "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable"

func main() {
	port, dbURL := loadEnv()

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

	log.Printf("Started server on port %s...\n", port)

	// Start http server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func loadEnv() (string, string) {
	godotenv.Load(".env")

	environment := os.Getenv("ENVIRONMENT")
	switch environment {
	case "development":
		return DEFAULT_PORT, DEVELOPMENT_DB_URL
	case "production":
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			log.Fatal("DB_URL is not found in the environment")
		}

		return port, dbURL

	default:
		log.Fatal("ENVIRONMENT is not found in the environment")
		return "", ""
	}
}
