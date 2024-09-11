package validation

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/go-playground/validator/v10"
)

// validate - глобальная переменная для экземпляра валидатора.
var validate *validator.Validate

// init инициализирует экземпляр валидатора при запуске пакета.
func init() {
	validate = validator.New()
}

// ValidateUser проверяет корректность данных пользователя с помощью валидатора.
func ValidateUser(user models.User) error {
	err := validate.Struct(user)
	if err != nil {
		return errors.New("invalid user data")
	}
	return nil
}

// ValidateBooking проверяет корректность данных бронирования с помощью валидатора.
func ValidateBooking(booking models.Booking) error {
	err := validate.Struct(booking)
	if err != nil {
		return errors.New("invalid booking data")
	}
	return nil
}
