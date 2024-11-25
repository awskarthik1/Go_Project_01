package main

import (
	"fmt"
	"log"

	config "example.com/kartheek/go_project02/Config"
)

func main() {
	// Define DSN to connect to the local MySQL database
	dsn := "root:@Lms230695@tcp(127.0.0.1:3306)/mydatabase"

	// Initialize Config
	cfg := config.NewConfig()

	// Connect to the database
	if err := cfg.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer cfg.DB.Close()
	fmt.Println("Connected to the MySQL database!")

	// Test database connection
	var version string
	err := cfg.DB.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to fetch database version: %v", err)
	}
	fmt.Printf("Database version: %s\n", version)

	// Make an HTTP GET request
	response, err := cfg.MakeRequest("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()
	fmt.Println("HTTP request succeeded with status:", response.Status)
}
