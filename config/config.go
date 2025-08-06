package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error Initialize SQLite: %v", err)
	}

	return nil
}

// Return Database Local the package config
func GetSQLite() *gorm.DB {
	return db
}


func GETLogger(p string) *Logger {
	// Initialize Logger

	logger = NewLogger(p)
	return logger
}
