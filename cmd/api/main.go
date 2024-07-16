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
const DEFAULT_WS_PORT = "5050"
const DEVELOPMENT_DB_URL = "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable"

func main() {
	port, dbURL, wsPort := loadEnv()

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

	// Start http server in a goroutine
	go func() {
		log.Printf("[HTTP] Started server on port %s...\n", port)

		// Start http server
		err = srv.ListenAndServe()
		if err != nil {
			log.Panic(err)
		}
	}()

	http.HandleFunc("/ws", webSocketHandler)
	log.Printf("[WS] Started WebSocket server on port %s...\n", wsPort)
	err = http.ListenAndServe(":"+wsPort, nil)
	if err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}
}

func loadEnv() (string, string, string) {
	godotenv.Load(".env")

	environment := os.Getenv("ENVIRONMENT")
	switch environment {
	case "development":
		return DEFAULT_PORT, DEVELOPMENT_DB_URL, DEFAULT_WS_PORT
	case "production":
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			log.Fatal("DB_URL is not found in the environment")
		}

		wsPort := os.Getenv("WS_PORT")
		if wsPort == "" {
			wsPort = "5050"
		}

		return port, dbURL, wsPort

	default:
		log.Fatal("ENVIRONMENT is not found in the environment")
		return "", "", ""
	}
}
