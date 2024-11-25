package config

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// Config struct holds database and HTTP configurations
type Config struct {
	DB     *sql.DB
	Client *http.Client
}

// NewConfig initializes and returns a Config struct
func NewConfig() *Config {
	return &Config{
		Client: &http.Client{},
	}
}

// Connect connects to the MySQL database
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
