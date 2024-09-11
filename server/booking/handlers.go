package booking

import (
	"encoding/json"
	"net/http"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/validation"
	"github.com/gorilla/mux"
)

// BookingHandlers представляет структуру для работы с обработчиками запросов, связанных с бронированием.
type BookingHandlers struct {
	BookingService service.BookingService
}

// NewBookingHandlers создаёт новый экземпляр BookingHandlers с предоставленной службой бронирования.
func NewBookingHandlers(bookingService service.BookingService) *BookingHandlers {
	return &BookingHandlers{BookingService: bookingService}
}

// CreateBooking обрабатывает запрос на создание нового бронирования.
// Декодирует тело запроса в структуру Booking, валидирует данные и сохраняет бронирование.
// Если успешен, возвращает созданное бронирование в формате JSON.

// @Summary Create a new booking
// @Description Create a new booking for a user, validating the data and saving it to the database.
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body models.Booking true "Booking data"
// @Success 201 {object} models.Booking
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /booking [post]
func (h *BookingHandlers) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking models.Booking
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Преобразуем время в UTC.
	newBooking.Start_time = newBooking.Start_time.UTC()
	newBooking.End_time = newBooking.End_time.UTC()

	// Валидируем данные бронирования.
	err = validation.ValidateBooking(newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Создаём бронирование.
	err = h.BookingService.Create(&newBooking)
	if err != nil {
		http.Error(w, `{"error":"error saving booking"}`, http.StatusInternalServerError)
		return
	}

	// Возвращаем созданное бронирование.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// DeleteBooking обрабатывает запрос на удаление бронирования по ID.
// Проверяет наличие ID и существование бронирования перед удалением.

// @Summary Delete a booking by ID
// @Description Delete an existing booking using its ID.
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Booking not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /booking/{id} [delete]
func (h *BookingHandlers) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из параметров URL.
	params := mux.Vars(r)
	id := params["id"]

	// Проверка на наличие ID в запросе.
	if id == "" {
		http.Error(w, `{"error":"id is not found"}`, http.StatusBadRequest)
		return
	}

	// Проверка на существование бронирования.
	_, err := h.BookingService.Get(id)
	if err != nil {
		http.Error(w, `{"error":"bookig is not found"}`, http.StatusBadRequest)
		return
	}

	// Удаление бронирования.
	err = h.BookingService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetBookings обрабатывает запрос на получение всех бронирований.
// Возвращает список бронирований в формате JSON.

// @Summary Get all bookings
// @Description Retrieve a list of all bookings.
// @Tags bookings
// @Accept json
// @Produce json
// @Success 200 {array} models.Booking "List of bookings"
// @Success 204 "No content"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bookings [get]
func (h *BookingHandlers) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.BookingService.GetAll()
	if err != nil {
		http.Error(w, "Error fetching bookings", http.StatusInternalServerError)
		return
	}

	// Если бронирований нет, возвращаем статус 204 (No Content).
	if len(bookings) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Возвращаем список бронирований.
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
