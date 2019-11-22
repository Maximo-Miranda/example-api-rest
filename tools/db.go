package tools

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect -> returns a new instance of gorm DB or an error if the connection fails
func Connect() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", getConnectionString())

	db.LogMode(true)

	return db, err
}

// getConnectionString -> return connection string for postgres
func getConnectionString() string {

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("EXAMPLE_API_REST_DB_HOST"),
		os.Getenv("EXAMPLE_API_REST_DB_PORT"),
		os.Getenv("EXAMPLE_API_REST_DB_USER"),
		os.Getenv("EXAMPLE_API_REST_DB_NAME"),
		os.Getenv("EXAMPLE_API_REST_DB_PASSWORD"),
	)
}



