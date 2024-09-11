package service

import "github.com/AramLab/booking-service/models"

// UserService описывает интерфейс для работы с пользователями.
// Интерфейс включает методы для создания, удаления, получения всех пользователей и поиска по ID.
type UserService interface {
	Create(user *models.User) error
	Delete(id string) error
	GetAll() ([]models.User, error)
	Get(id string) (models.User, error)
}

// BookingService описывает интерфейс для работы с бронированиями.
// Интерфейс включает методы для создания, удаления, получения всех бронирований и поиска по ID.
type BookingService interface {
	Create(booking *models.Booking) error
	Delete(id string) error
	GetAll() ([]models.Booking, error)
	Get(id string) (models.Booking, error)
}

// Service представляет общую структуру, которая объединяет интерфейсы для работы с пользователями и бронированиями.
// Содержит экземпляры сервисов для пользователей и бронирований.
type Service struct {
	UserServ    UserService
	BookingServ BookingService
}
