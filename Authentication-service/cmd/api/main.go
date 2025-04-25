package main

import (
	"Authentication-service/data"
	"database/sql"
	"log"
)

const PORT = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication service")

}
