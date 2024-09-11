package user

import (
	"net/http"

	"github.com/AramLab/booking-service/service"
)

type UserHandlers struct {
	UserService service.UserService
}

func NewUserHandlers(userService service.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
