package booking

import (
	"encoding/json"
	"net/http"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/validation"
	"github.com/gorilla/mux"
)

type BookingHandlers struct {
	BookingService service.BookingService
}

func NewBookingHandlers(bookingService service.BookingService) *BookingHandlers {
	return &BookingHandlers{BookingService: bookingService}
}

func (h *BookingHandlers) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking models.Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBooking.Start_time = newBooking.Start_time.UTC()
	newBooking.End_time = newBooking.End_time.UTC()

	err = validation.ValidateBooking(newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.BookingService.Create(&newBooking)
	if err != nil {
		http.Error(w, `{"error":"error saving booking"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *BookingHandlers) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	//
	if id == "" {
		http.Error(w, `{"error":"id is not found"}`, http.StatusBadRequest)
		return
	}

	_, err := h.BookingService.Get(id)
	if err != nil {
		http.Error(w, `{"error":"bookig is not found"}`, http.StatusBadRequest)
		return
	}
	//
	err = h.BookingService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *BookingHandlers) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.BookingService.GetAll()
	if err != nil {
		http.Error(w, "Error fetching bookings", http.StatusInternalServerError)
		return
	}

	if len(bookings) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
