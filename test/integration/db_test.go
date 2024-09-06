package db_test

import (
	"integrationtests/internal/adapter/repository/postgre/db"
	"testing"
)

func Connected() bool {
	// Implement the logic to check if connected to the database
	err := db.SqlDB.Ping()
	return err == nil
}

func TestConnect(t *testing.T) {
	// Connect to the database
	db.Init()

	// Check if connected to the database
	if Connected() {
		t.Log("Connected to the database")
	} else {
		t.Error("Failed to connect to the database")
	}

	// Close the database connection
	db.Close()
}
