package main

import (
	handler "integrationtests/internal/adapter/handler/http"
	repository "integrationtests/internal/adapter/repository/mysql"
	"integrationtests/internal/adapter/repository/mysql/db"
	"integrationtests/internal/usecase/service"
)

func main() {

	// Initialize database connection
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	//Initialize repository
	ur := repository.NewUserRepository(db)

	// Initialize service
	us := service.NewUserService(ur)

	uh := &handler.UserHandler{
		// initialize any dependencies here
		UserService: us,
	}

	uh.Init()
}
