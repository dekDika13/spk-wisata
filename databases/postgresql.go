package databases

import (
	"backend/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Address  string
	DB_Name     string
	DB_Port     string // Added port for PostgreSQL
}

func InitDB() {
	// Load configuration from environment variables
	config := Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Address:  os.Getenv("DB_ADDRESS"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Port:     os.Getenv("DB_PORT"), // Ensure this is set in the environment
	}

	// Create connection string for PostgreSQL
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_Address, config.DB_Username, config.DB_Password, config.DB_Name, config.DB_Port)

	var err error
	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Run migrations
	InitMigrate()
}

func InitMigrate() {
	// Migrate models
	_ = DB.AutoMigrate(
		
		&model.User{},
		&model.Criteria{},
		&model.Destination{},
		&model.Review{},
		&model.DetailCriteria{},
		&model.Profile{},
	)
}
