package config

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	_ "modernc.org/sqlite" // Go SQLite driver for Windows
)

var DB *gorm.DB

// ConnectDatabase initializes the DB connection based on DB_TYPE
func ConnectDatabase() {
	dbType := os.Getenv("DB_TYPE") // postgres, mysql, sqlite, sqlserver

	var dsn string
	var dialector gorm.Dialector

	switch dbType {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		dialector = mysql.Open(dsn)

	case "sqlite":
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			log.Fatal("DB_NAME environment variable is required for SQLite")
		}

		if runtime.GOOS == "windows" {
			// Use Go driver on Windows to avoid CGO issues
			database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
			if err != nil {
				log.Fatal("Failed to connect to SQLite (Windows):", err)
			}
			DB = database
		} else {
			// For other OSes, normal SQLite driver
			dialector = sqlite.Open(dbName)
		}

	case "sqlserver":
		dsn = fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		dialector = sqlserver.Open(dsn)

	default:
		log.Fatalf("Unsupported DB_TYPE: %s", dbType)
	}

	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
}
