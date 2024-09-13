package booking

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AramLab/booking-service/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockBookingService struct {
	mock.Mock
}

func (m *MockBookingService) Create(booking *models.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookingService) GetAll() ([]models.Booking, error) {
	args := m.Called()
	return args.Get(0).([]models.Booking), args.Error(1)
}

func (m *MockBookingService) Get(id string) (*models.Booking, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Booking), args.Error(1)
}

func TestCreateBooking(t *testing.T) {
	mockService := new(MockBookingService)
	handler := NewBookingHandlers(mockService)

	newBooking := models.Booking{
		ID:         1,
		User_id:    123,
		Start_time: time.Now().UTC(),
		End_time:   time.Now().Add(time.Hour).UTC(),
	}

	mockService.On("Create", &newBooking).Return(nil)

	body, err := json.Marshal(newBooking)
	if err != nil {
		t.Errorf("error to use func marshal")
	}
	req := httptest.NewRequest("POST", "/booking", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateBooking(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusCreated, res.StatusCode)

	var createdBooking models.Booking
	err = json.NewDecoder(res.Body).Decode(&createdBooking)
	require.NoError(t, err)

	require.Equal(t, newBooking.ID, createdBooking.ID)
	require.Equal(t, newBooking.User_id, createdBooking.User_id)
}

func TestDeleteBooking(t *testing.T) {
	mockService := new(MockBookingService)
	handler := NewBookingHandlers(mockService)

	mockService.On("Get", "1").Return(&models.Booking{ID: 1}, nil)
	mockService.On("Delete", "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/booking/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handler.DeleteBooking(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetBookings(t *testing.T) {
	mockService := new(MockBookingService)
	handler := NewBookingHandlers(mockService)

	bookings := []models.Booking{
		{ID: 1, User_id: 123, Start_time: time.Now(), End_time: time.Now().Add(time.Hour)},
		{ID: 2, User_id: 456, Start_time: time.Now(), End_time: time.Now().Add(2 * time.Hour)},
	}

	mockService.On("GetAll").Return(bookings, nil)

	req := httptest.NewRequest("GET", "/bookings", nil)
	w := httptest.NewRecorder()

	handler.GetBookings(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var fetchedBookings []models.Booking
	err := json.NewDecoder(res.Body).Decode(&fetchedBookings)
	require.NoError(t, err)

	require.Equal(t, 2, len(fetchedBookings))
	require.Equal(t, bookings[0].ID, fetchedBookings[0].ID)
	require.Equal(t, bookings[1].ID, fetchedBookings[1].ID)
}

func TestGetBookings_NoContent(t *testing.T) {
	mockService := new(MockBookingService)
	handler := NewBookingHandlers(mockService)

	mockService.On("GetAll").Return([]models.Booking{}, nil)

	req := httptest.NewRequest("GET", "/bookings", nil)
	w := httptest.NewRecorder()

	handler.GetBookings(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCreatBooking_InvalidData(t *testing.T) {
	mockService := new(MockBookingService)
	handler := NewBookingHandlers(mockService)

	booking := models.Booking{
		ID: 1,
	}

	mockService.On("Create", &booking).Return(errors.New("error of creating booking"))

	body, err := json.Marshal(booking)
	if err != nil {
		t.Errorf("error to use func marshal")
	}
	req := httptest.NewRequest("POST", "/booking", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateBooking(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}
