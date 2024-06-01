package main

import (
	"fmt"
	"log"

	"github.com/Rohan2998/the-dating-app-backend/config"
	"github.com/Rohan2998/the-dating-app-backend/pkg/database"
)

func main() {

	fmt.Println("Welcome to the dating app")

	// Load configuration
	ptrConfig := config.LoadConfig()

	// Connect to the database
	db, err := database.Connect(ptrConfig)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connection to database successfully done, starting server now..")

}
