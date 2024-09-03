package main

import (
	handler "integrationtests/internal/adapter/handler/http"
	repository "integrationtests/internal/adapter/repository/postgre"
	"integrationtests/internal/adapter/repository/postgre/db"
	"integrationtests/internal/usecase/service"
)

func main() {

	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	//Initialize repository
	ur := repository.NewRepoUser(db)

	// Initialize service
	us := service.NewUserService(ur)

	uh := &handler.UserHandler{
		// initialize any dependencies here
		UserService: *us,
	}

	uh.Init()
}
