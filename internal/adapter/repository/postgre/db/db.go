// write data base connection here
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

// Connect connects to the database
func Init() (*sql.DB, error) {
	var err error

	// Load configuration from YAML file
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	DBCon, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		"?parseTime=true",
	))
	if err != nil {
		log.Fatal(err)
	}

	// Retry mechanism
	retries := 3
	for retries > 0 {
		err = DBCon.Ping()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to the database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		retries--
	}

	if err != nil {
		log.Fatal(err)
	}

	// Set database connection pool parameters
	DBCon.SetMaxOpenConns(25)
	DBCon.SetMaxIdleConns(25)
	DBCon.SetConnMaxLifetime(5 * 60) // 5 minutes

	return DBCon, err
}

// loadConfig loads the configuration from a YAML file
func loadConfig() (*Config, error) {
	// Open the YAML file
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Parse the YAML file into a Config struct
	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// Config represents the database configuration
type Config struct {
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
}

// Close closes the database connection
func Close() {
	DBCon.Close()
}
