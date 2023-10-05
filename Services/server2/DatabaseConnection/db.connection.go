package databaseconnectionServer2

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadEnvServer2() {
	envFileNameServer2 := ".env.server2" // Specify the file name here
	if err := godotenv.Load(envFileNameServer2); err != nil {
		log.Fatalf("Error loading %s file: %v", envFileNameServer2, err)
	}
}
func configValues() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)

	return dbURI
}

func Databaseconnectionserver2() {
	loadEnvServer2()
	dburi := configValues()
	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		// handle error
		log.Println(`Could not get SQL DB`, err)
	}
	sqlDB.Close()
}

// study dusty
