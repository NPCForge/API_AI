package config

import (
	"bufio"
	"database/sql"
	"fmt"
	. "my-api/pkg"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbClient *sql.DB
	once     sync.Once
)

// IsRunningInDocker checks if the application is running inside a Docker container.
func IsRunningInDocker() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "/docker/") || strings.Contains(scanner.Text(), "/kubepods/") {
			return true
		}
	}
	return false
}

// InitDB initializes the PostgreSQL database connection with retry logic.
func InitDB() {
	once.Do(func() {
		var err error

		// Database connection configuration
		connStr := "user=" + GetEnvVariable("POSTGRES_USER")
		connStr += " password=" + GetEnvVariable("POSTGRES_PASSWORD")
		connStr += " dbname=" + GetEnvVariable("POSTGRES_DB")

		if IsRunningInDocker() {
			connStr += " host=postgres"
		} else {
			connStr += " host=" + GetEnvVariable("POSTGRES_HOST")
		}

		connStr += " port=" + GetEnvVariable("POSTGRES_PORT") + " sslmode=disable"
		maxRetries := 4
		retryDelay := 10 * time.Second

		for retries := 0; retries < maxRetries; retries++ {
			// Try to open the connection
			dbClient, err = sql.Open("postgres", connStr)
			if err != nil {
				msg := fmt.Sprintf("Error opening connection to PostgreSQL (attempt %d/%d)", retries+1, maxRetries)
				DisplayContext(msg, Error, err)
			} else {
				// Test the connection
				err = dbClient.Ping()
				if err == nil {
					DisplayContext("Successfully connected to PostgreSQL!", Update)
					return
				}
				msg := fmt.Sprintf("Connection error to PostgreSQL (attempt %d/%d)", retries+1, maxRetries)
				DisplayContext(msg, Error, err)
			}

			// Close the connection on error to avoid leaks
			if dbClient != nil {
				dbClient.Close()
			}

			// Wait before retrying
			if retries < maxRetries-1 {
				DisplayContext(fmt.Sprintf("Retrying in %v", retryDelay), Update)
				time.Sleep(retryDelay)
			}
		}

		// If all attempts fail
		msg := fmt.Sprintf("Failed to connect to PostgreSQL after %d attempts: %v", maxRetries, err)
		DisplayContext(msg, Error, true)
	})
}

// GetDB returns the active PostgreSQL database client.
func GetDB() *sql.DB {
	if dbClient == nil {
		InitDB()
	}
	return dbClient
}

// CloseDB closes the PostgreSQL database connection.
func CloseDB() {
	if dbClient != nil {
		err := dbClient.Close()
		if err != nil {
			DisplayContext("Error closing PostgreSQL connection", Error, err)
		} else {
			fmt.Println("PostgreSQL connection closed.")
		}
	}
}
