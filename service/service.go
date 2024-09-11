package service

import "github.com/AramLab/booking-service/models"

type UserService interface {
	Create(user *models.User) error
	Delete(id string) error
	GetAll() ([]models.User, error)
	Get(id string) (models.User, error)
}

type BookingService interface {
	Create(booking *models.Booking) error
	Delete(id string) error
	GetAll() ([]models.Booking, error)
	Get(id string) (models.Booking, error)
}

type Service struct {
	UserServ    UserService
	BookingServ BookingService
}
