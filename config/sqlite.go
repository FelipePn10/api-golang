package config

import (
	"github.com/FelipePn10/api-golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.go"
	// Check if the database file exits
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file not found! Creating new database file.")
		// Create the database file and directory
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening SQLite database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error auto-migrating SQLite database: %v", err)
		return nil, err
	}
	return db, nil
}
