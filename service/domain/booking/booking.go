package booking

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/storage"
)

type BookingService struct {
	BookingRepository storage.BookingRepository
}

func NewBookingService(BookingRepository storage.BookingRepository) *BookingService {
	return &BookingService{BookingRepository: BookingRepository}
}

func (bs *BookingService) Create(booking *models.Booking) error {
	err := bs.BookingRepository.Save(booking)
	if err != nil {
		return errors.New(`{"error":"error of creating booking"}`)
	}
	return nil
}

func (bs *BookingService) Delete(id string) error {
	if id == "" {
		return errors.New(`{"error":"id is not specified"}`)
	}
	_, err := bs.Get(id)
	if err != nil {
		return errors.New(`{"error":"booking is not found"}`)
	}

	err = bs.BookingRepository.DeleteById(id)
	if err != nil {
		return errors.New(`{"error":"error to delete booking"}`)
	}
	return nil
}

func (bs *BookingService) Get(id string) (models.Booking, error) {
	user, err := bs.BookingRepository.FindById(id)
	if err != nil {
		return models.Booking{}, errors.New(`{"error":"error to get booking"}`)
	}
	return user, nil
}

func (bs *BookingService) GetAll() ([]models.Booking, error) {
	bookings, err := bs.BookingRepository.FindAll()
	if err != nil {
		return []models.Booking{}, errors.New(`{"error":"error get all bookings"}`)
	}
	return bookings, nil
}
