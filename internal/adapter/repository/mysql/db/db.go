package db

import (
	"database/sql"
	"fmt"
	"go-sample/internal/domain/model"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
	SqlDB *sql.DB
)

func Init() (*gorm.DB, error) {
	// Load configuration from YAML file
	config, err := loadConfig("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Build connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	// Open connection
	DBCon, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}

	sqlDB, err := DBCon.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sql.DB from gorm.DB: %w", err)
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
		return nil, fmt.Errorf("failed to ping database after retries: %w", err)
	}

	// Set database connection pool parameters
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DBCon.AutoMigrate(model.User{})

	return DBCon, nil
}

// loadConfig loads the configuration from a YAML file
func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}

func Close() {
	if SqlDB != nil {
		if err := SqlDB.Close(); err != nil {
			log.Printf("Error closing the database: %v", err)
		}
	}
}

type Config struct {
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
}
