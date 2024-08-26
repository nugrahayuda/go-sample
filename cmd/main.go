package main

import (
	"integrationtests/internal/db"
	"integrationtests/internal/repo"
	"integrationtests/internal/router"
	"integrationtests/internal/service"
)

func main() {

	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	//Initialize repository
	repoUser := repo.NewRepoUser(db)

	// Initialize service
	userService := service.NewUserService(repoUser)

	handler := &router.Service{
        // initialize any dependencies here
		UserService: userService,
    }

	handler.Init()
}
