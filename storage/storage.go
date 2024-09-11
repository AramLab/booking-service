package storage

import (
	"github.com/AramLab/booking-service/models"
)

// UserRepository определяет интерфейс для работы с пользователями в хранилище данных.
type UserRepository interface {
	Save(user *models.User) error
	DeleteById(id string) error
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
}

// BookingRepository определяет интерфейс для работы с бронированиями в хранилище данных.
type BookingRepository interface {
	Save(booking *models.Booking) error
	DeleteById(id string) error
	FindAll() ([]models.Booking, error)
	FindById(id string) (models.Booking, error)
}

// Repository объединяет интерфейсы для работы с пользователями и бронированиями,
// обеспечивая доступ к обоим типам репозиториев через единое хранилище.
type Repository struct {
	UserRepo    UserRepository
	BookingRepo BookingRepository
}
