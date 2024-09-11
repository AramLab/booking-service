package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/validation"
	"github.com/gorilla/mux"
)

type UserHandlers struct {
	UserService service.UserService
}

func NewUserHandlers(userService service.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validation.ValidateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = newUser.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, `{"error":"error hashing password"}`, http.StatusInternalServerError)
		return
	}

	newUser.Created_at = time.Now()
	newUser.Updated_at = time.Now()

	err = h.UserService.Create(&newUser)
	if err != nil {
		http.Error(w, `{"error":"error creating user"}`, http.StatusInternalServerError)
		return
	}

	// Убираем пароль из ответа
	newUser.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		http.Error(w, `{"error":"id is not found"}`, http.StatusBadRequest)
		return
	}

	_, err := h.UserService.Get(id)
	if err != nil {
		http.Error(w, `{"error":"user is not found"}`, http.StatusNotFound)
		return
	}

	err = h.UserService.Delete(id)
	if err != nil {
		http.Error(w, `{"error":"error to delete user"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
