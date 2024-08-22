package main

import (
	"integrationtests/internal/db"
	"integrationtests/internal/repo"
	"integrationtests/internal/router"
)

func main() {
	router.Init()
	db, err := db.Init()
	if err!= nil {
        panic(err)
    }

	repo.NewRepoUser(db)
}