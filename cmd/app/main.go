package main

import (
	handler "integrationtests/internal/adapter/handler/http"
	"integrationtests/internal/adapter/repository/postgre/db"
	"integrationtests/internal/adapter/repository/postgre"
	"integrationtests/internal/usecase/service"
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

	handler := &handler.Service{
		// initialize any dependencies here
		UserService: userService,
	}

	handler.Init()
}
