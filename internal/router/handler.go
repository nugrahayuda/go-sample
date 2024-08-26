package router

import (
	"fmt"
	"integrationtests/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	UserService service.UserService
}

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

func (s *Service) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// handle user get request
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	data, e := s.UserService.GetUserByID(id)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting user: %v\n", e)
		return
	}

	fmt.Fprintf(w, "User: %+v\n", data)
}
