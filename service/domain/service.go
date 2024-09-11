package domain

import (
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/service/domain/booking"
	"github.com/AramLab/booking-service/service/domain/user"
	"github.com/AramLab/booking-service/storage"
)

func NewService(repo *storage.Repository) *service.Service {
	return &service.Service{
		UserServ:    user.NewUserService(repo.UserRepo),
		BookingServ: booking.NewBookingService(repo.BookingRepo),
	}
}
