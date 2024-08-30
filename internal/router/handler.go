package router

import (
	"encoding/json"
	"fmt"
	"integrationtests/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	UserService service.UserService
}

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	UserService service.UserService
}

func (s Service) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// handle user get request
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	data, e := s.UserService.GetUserByID(id)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error getting user: %v", e)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
