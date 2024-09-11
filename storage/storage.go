package storage

import (
	"github.com/AramLab/booking-service/models"
)

type UserRepository interface {
	Save(user *models.User) error
	DeleteById(id string) error
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
}

type BookingRepository interface {
	Save(booking *models.Booking) error
	DeleteById(id string) error
	FindAll() ([]models.Booking, error)
	FindById(id string) (models.Booking, error)
}

type Repository struct {
	UserRepo    UserRepository
	BookingRepo BookingRepository
}
