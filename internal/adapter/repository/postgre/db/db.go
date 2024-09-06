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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
	SqlDB *sql.DB
)

// Connect connects to the database
func Init() (*gorm.DB, error) {
	var err error

	// Load configuration from YAML file
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	DBCon, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := DBCon.DB()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Retry mechanism
	retries := 3
	for retries > 0 {

		err = sqlDB.Ping()
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

	// // Set database connection pool parameters
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * 60) // 5 minutes

	return DBCon, err
}

// loadConfig loads the configuration from a YAML file
func loadConfig() (*Config, error) {
	// Open the YAML file
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatal("Fatal config", err)
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
	SqlDB.Close()
}
