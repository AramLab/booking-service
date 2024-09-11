package postgres

import (
	"github.com/AramLab/booking-service/storage"
	"github.com/AramLab/booking-service/storage/postgres/booking"
	"github.com/AramLab/booking-service/storage/postgres/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewRepository(db *pgxpool.Pool) *storage.Repository {
	return &storage.Repository{
		UserRepo:    user.NewUserRepository(db),
		BookingRepo: booking.NewBookingRepository(db),
	}
}
