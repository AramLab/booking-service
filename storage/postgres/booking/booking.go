package booking

import (
	"context"
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{db: db}
}

func (br *BookingRepository) Save(booking *models.Booking) error {
	_, err := br.db.Exec(context.Background(), "INSERT INTO booking (user_id, start_time, end_time) VALUES ($1, $2, $3)",
		booking.User_id, booking.Start_time, booking.End_time)
	if err != nil {
		return errors.New(`{"error":"failed to add booking"}`)
	}
	return nil
}

func (br *BookingRepository) DeleteById(id string) error {
	_, err := br.db.Exec(context.Background(), "DELETE FROM booking WHERE id=$1", id)
	if err != nil {
		return errors.New(`{"error":"failed to delete booking"}`)
	}
	return nil
}

func (br *BookingRepository) FindAll() ([]models.Booking, error) {
	rows, err := br.db.Query(context.Background(), "SELECT id, user_id, start_time, end_time FROM bookings")
	if err != nil {
		return []models.Booking{}, errors.New(`{"error":"failed to find all bookings"}`)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(&booking.ID, &booking.User_id, &booking.Start_time, &booking.End_time)
		if err != nil {
			return []models.Booking{}, errors.New(`{"error":"error scanning bookings"}`)
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (br *BookingRepository) FindById(id string) (models.Booking, error) {
	var booking models.Booking
	row := br.db.QueryRow(context.Background(), "SELECT id, user_id, start_time, end_time FROM booking WHERE id=$1", id)
	err := row.Scan(&booking.ID, &booking.User_id, &booking.Start_time, &booking.End_time)
	if err != nil {
		return models.Booking{}, errors.New(`{"error":"error scanning booking"}`)
	}
	return booking, nil
}
