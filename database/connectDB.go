package database

// dropped support for database & docker

/*
import (
	"fmt"
	"os"

	"goscrape.com/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {

	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")
}
*/
