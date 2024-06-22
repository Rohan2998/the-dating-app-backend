package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rohan2998/the-dating-app-backend/config"
	"github.com/Rohan2998/the-dating-app-backend/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run cmd/db/migrations.go [up|down]")
		return
	}

	fmt.Println("initializing migrations..")

	// load configuration
	ptrConfig := config.LoadConfig()

	// connect to the database
	db, err := db.Connect(ptrConfig)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	defer db.Close()

	// Create a new migration driver for MySQL
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not create migration driver: %v", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/xampp/htdocs/the-dating-app/backend/db/migrations", // Path to your migration files
		"mysql", driver)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	// get second argument from command line and execute accordingly
	action := args[1]

	// switch case for action
	switch action {
	case "up":
		fmt.Println("running up migrations..")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not apply migrations: %v", err)
		}
		fmt.Println("migrations applied successfully")
	case "down":
		fmt.Println("running down migrations..")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("could not rollback migrations: %v", err)
		}
		fmt.Println("migrations rolled back successfully")
	default:
		fmt.Println("Unknown command. Usage: go run cmd/db/migrations.go [up|down]")
	}
}
