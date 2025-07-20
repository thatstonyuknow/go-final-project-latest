package db

import (
	"database/sql"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

type DbConfig struct {
	DbFile   string  // Database file name
	DbDriver string  // Database driver name
	dbConn   *sql.DB // Database connection
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		DbFile:   getEnvWithDefault("TODO_DB_FILE", "scheduler.db"),
		DbDriver: "sqlite",
		dbConn:   nil,
	}
}

const (
	// SQL commands
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS scheduler (
			id      INTEGER PRIMARY KEY AUTOINCREMENT,        -- auto-incrementing identifier
			date    TEXT    NOT NULL,                         -- date in format 'YYYYMMDD' (or Go format '20060102')
			title   TEXT    NOT NULL,                         -- task title
			comment TEXT,                                     -- comment for the task
			repeat  TEXT    NOT NULL CHECK(LENGTH(repeat) <= 128)  --
		);
	`

	createIndexSQL = `
		CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
	`
)

func InitDB() (*sql.DB, error) {
	// Create default database configuration
	config := NewDbConfig()
	return InitDBWithConfig(config)
}

func InitDBWithConfig(config *DbConfig) (*sql.DB, error) {
	// Check if the database file is specified
	if config.DbFile == "" {
		log.Error("Database name is empty")
		return nil, errors.New("database name is empty")
	}

	// Check if the database file exists
	_, err := os.Stat(config.DbFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infof("Database file %s does not exist", config.DbFile)
			log.Infof("Creating new database file: %s", config.DbFile)
			// Create the database file
			file, err := os.Create(config.DbFile)
			if err != nil {
				log.Error("Error creating database file: ", err)
				return nil, err
			}
			defer file.Close()
		}
	}

	// setting up the database connection
	dbConn, err := sql.Open(config.DbDriver, config.DbFile)
	if err != nil {
		log.Error("Error opening DB: ", err)
		return nil, err
	}

	// Check if the database connection is valid
	err = dbConn.Ping()
	if err != nil {
		log.Error("Error pinging DB: ", err)
		dbConn.Close() // Clean up on error
		return nil, err
	}

	// Create the scheduler table if it does not exist
	log.Info("Creating scheduler table if it does not exist")
	_, err = dbConn.Exec(createTableSQL)
	if err != nil {
		log.Error("Error creating scheduler table: ", err)
		dbConn.Close() // Clean up on error
		return nil, err
	}

	// Create indexes for the scheduler table
	log.Info("Creating indexes for the scheduler table")
	_, err = dbConn.Exec(createIndexSQL)
	if err != nil {
		log.Error("Error creating indexes for the scheduler table: ", err)
		dbConn.Close() // Clean up on error
		return nil, err
	}

	// Store connection in config for potential future use
	config.dbConn = dbConn

	log.Info("Database initialized successfully")

	return dbConn, nil
}

func getEnvWithDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
