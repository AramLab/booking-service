package domain

import (
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/service/domain/booking"
	"github.com/AramLab/booking-service/service/domain/user"
	"github.com/AramLab/booking-service/storage"
)

// NewService создает новый экземпляр структуры Service, которая объединяет различные сервисы приложения
// для работы с пользователями и бронированиями. Входящим параметром является указатель на структуру Repository,
// которая предоставляет доступ к репозиториям пользователей и бронирований.
func NewService(repo *storage.Repository) *service.Service {
	return &service.Service{
		UserServ:    user.NewUserService(repo.UserRepo),
		BookingServ: booking.NewBookingService(repo.BookingRepo),
	}
}
