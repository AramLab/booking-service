package booking

import (
	"context"
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// BookingRepository представляет собой структуру для работы с сущностью бронирования в базе данных.
type BookingRepository struct {
	db *pgxpool.Pool
}

// NewBookingRepository создаёт новый репозиторий для работы с бронированиями, принимая пул соединений с базой данных.
func NewBookingRepository(db *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{db: db}
}

// Save сохраняет новую запись о бронировании в базу данных.
// При успешном добавлении возвращает ID нового бронирования.
func (br *BookingRepository) Save(booking *models.Booking) error {
	err := br.db.QueryRow(context.Background(),
		"INSERT INTO booking (user_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING id",
		booking.User_id, booking.Start_time, booking.End_time).Scan(&booking.ID)
	if err != nil {
		return errors.New(`{"error":"failed to add booking"}`)
	}
	return nil
}

// DeleteById удаляет запись о бронировании по его ID.
func (br *BookingRepository) DeleteById(id string) error {
	_, err := br.db.Exec(context.Background(), "DELETE FROM booking WHERE id=$1", id)
	if err != nil {
		return errors.New(`{"error":"failed to delete booking"}`)
	}
	return nil
}

// FindAll возвращает список всех бронирований из базы данных.
func (br *BookingRepository) FindAll() ([]models.Booking, error) {
	rows, err := br.db.Query(context.Background(), "SELECT id, user_id, start_time, end_time FROM booking")
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

// FindById ищет запись о бронировании по его ID.
func (br *BookingRepository) FindById(id string) (models.Booking, error) {
	var booking models.Booking
	row := br.db.QueryRow(context.Background(), "SELECT id, user_id, start_time, end_time FROM booking WHERE id=$1", id)
	err := row.Scan(&booking.ID, &booking.User_id, &booking.Start_time, &booking.End_time)
	if err != nil {
		return models.Booking{}, errors.New(`{"error":"error scanning booking"}`)
	}
	return booking, nil
}
