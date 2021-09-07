package main

import (
	"log"
	"os"

	"github.com/ariffinsetya/micro/flash-auth/pkg/database"
)

func main() {
	address := os.Getenv("AUTH_REDIS_ADDRESS")
	password := os.Getenv(("AUTH_REDIS_PASSWORD"))
	database, err := database.NewDatabase(address, password)
	if err != nil {
		log.Fatalf("Databse not available")
	}
}
