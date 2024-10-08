package main

import (
	handler "go-sample/internal/adapter/handler/http"
	repository "go-sample/internal/adapter/repository/mysql"
	"go-sample/internal/adapter/repository/mysql/db"
	"go-sample/internal/usecase/service"
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
