package handler

import (
	"encoding/json"
	"fmt"
	"integrationtests/internal/domain/model"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService UserServiceInterface
}

type UserServiceInterface interface {
	GetUserByID(id string) (model.User, error)
}

func (h UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// handle user get request
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	data, e := h.UserService.GetUserByID(id)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Error getting user: %v", e)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
