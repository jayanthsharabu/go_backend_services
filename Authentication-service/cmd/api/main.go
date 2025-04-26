package main

import (
	"Authentication-service/data"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const PORT = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication service")

	conn, err := connectToDB()
	if err != nil {
		log.Panic("Cannot connect to database:", err)
	}

	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// Start the server in a goroutine
	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	conn.Close()

	log.Println("Server exited properly")
}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, errors.New("DSN environment variable not set")
	}

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection, nil
		}

		if counts > 10 {
			log.Println(err)
			return nil, errors.New("failed to connect to Postgres")
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
