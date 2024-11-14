package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Config struct holds database and HTTP configurations
type Config struct {
	DB     *sql.DB
	Client *http.Client
}

// Connect to the MySQL database
func (config *Config) Connect(dsn string) error {
	var err error
	config.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return config.DB.Ping()
}

// MakeRequest sends a GET request to the specified URL
func (config *Config) MakeRequest(url string) (*http.Response, error) {
	return config.Client.Get(url)
}

func main() {
	// Define DSN to connect to the local MySQL database
	dsn := "root:@Lms230695@tcp(127.0.0.1:3306)/mydatabase"

	// Initialize Config with HTTP client
	config := Config{
		Client: &http.Client{},
	}

	// Connect to the database
	if err := config.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer config.DB.Close()
	fmt.Println("Connected to the MySQL database!")

	// test database connection
	var version string
	err := config.DB.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to fetch database version: %v", err)
	}
	fmt.Printf("Database version: %s\n", version)

	// Make an HTTP GET request
	response, err := config.MakeRequest("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()
	fmt.Println("HTTP request succeeded with status:", response.Status)
}
