package logger

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// This is the logger instance of sqlite3 DB. I chose it as a way to save the IPS and to which URLS are these requests going.
// The database is saved in the same directory as the main executable.
// The database is created when the application is run.
// We hold our Logs into a struct, we init the db, automigrate etc and we have a result variable that holds the requestLog data.

// Define a struct to represent the request log
type RequestLog struct {
	ID        uint `gorm:"primaryKey"`
	IP        string
	URL       string
	Timestamp time.Time
}

var db *gorm.DB

// InitializeDB initializes the database connection for logging
func InitializeDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("request-log.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// Migrate the RequestLog model to create the corresponding table
	err = db.AutoMigrate(&RequestLog{})
	if err != nil {
		return err
	}

	log.Println("💽 Connected to Log DB")
	return nil
}

// CloseDB closes the database connection for logging
func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// LogRequest saves the request logs to the database
func LogRequest(ip, url string) {
	// Create a new RequestLog instance
	requestLog := RequestLog{
		IP:        ip,
		URL:       url,
		Timestamp: time.Now(),
	}

	// Save the request log to the database
	result := db.Create(&requestLog)
	if result.Error != nil {
		log.Println("Failed to save request log:", result.Error)
	}
}
