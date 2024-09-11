package booking

import (
	"net/http"

	"github.com/AramLab/booking-service/service"
)

type BookingHandlers struct {
	BookingService service.BookingService
}

func NewBookingHandlers(bookingService service.BookingService) *BookingHandlers {
	return &BookingHandlers{BookingService: bookingService}
}

func (h *BookingHandlers) CreateBooking(w http.ResponseWriter, r *http.Request) {

}

func (h *BookingHandlers) DeleteBooking(w http.ResponseWriter, r *http.Request) {

}

func (h *BookingHandlers) GetBookings(w http.ResponseWriter, r *http.Request) {

}
