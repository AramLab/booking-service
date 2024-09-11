package validation

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user models.User) error {
	err := validate.Struct(user)
	if err != nil {
		return errors.New("invalid user data")
	}
	return nil
}

func ValidateBooking(booking models.Booking) error {
	err := validate.Struct(booking)
	if err != nil {
		return errors.New("invalid booking data")
	}
	return nil
}
