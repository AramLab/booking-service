package booking

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/storage"
)

// BookingService представляет собой слой бизнес-логики для управления бронированиями.
// Он взаимодействует с BookingRepository для выполнения операций над сущностями бронирования.
type BookingService struct {
	BookingRepository storage.BookingRepository
}

// NewBookingService создает и возвращает новый экземпляр BookingService с предоставленным репозиторием BookingRepository.
func NewBookingService(BookingRepository storage.BookingRepository) *BookingService {
	return &BookingService{BookingRepository: BookingRepository}
}

// Create создает новое бронирование, вызывая метод Save репозитория.
// Если происходит ошибка во время сохранения бронирования, возвращает соответствующую ошибку.
func (bs *BookingService) Create(booking *models.Booking) error {
	err := bs.BookingRepository.Save(booking)
	if err != nil {
		return errors.New(`{"error":"error of creating booking"}`)
	}
	return nil
}

// Delete удаляет бронирование по указанному идентификатору.
// Сначала проверяет, существует ли бронирование, затем вызывает метод DeleteById репозитория.
// Если бронирование не найдено или возникает ошибка при удалении, возвращает соответствующую ошибку.
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

// Get возвращает бронирование по указанному идентификатору, используя метод FindById репозитория.
// В случае ошибки возвращает пустое бронирование и ошибку.
func (bs *BookingService) Get(id string) (models.Booking, error) {
	user, err := bs.BookingRepository.FindById(id)
	if err != nil {
		return models.Booking{}, errors.New(`{"error":"error to get booking"}`)
	}
	return user, nil
}

// GetAll возвращает все существующие бронирования, используя метод FindAll репозитория.
// В случае ошибки возвращает пустой список и соответствующую ошибку.
func (bs *BookingService) GetAll() ([]models.Booking, error) {
	bookings, err := bs.BookingRepository.FindAll()
	if err != nil {
		return []models.Booking{}, errors.New(`{"error":"error get all bookings"}`)
	}
	return bookings, nil
}
