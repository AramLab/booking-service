package routes

import (
	"net/http"

	"github.com/AramLab/booking-service/server/booking"
	"github.com/AramLab/booking-service/server/user"
	"github.com/AramLab/booking-service/service"
	"github.com/gorilla/mux"
)

func RegisterRoutes(userService service.UserService, bookingService service.BookingService) http.Handler {
	r := mux.NewRouter()

	userHandler := user.NewUserHandlers(userService)
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")

	bookingHandler := booking.NewBookingHandlers(bookingService)
	r.HandleFunc("/booking", bookingHandler.CreateBooking).Methods("POST")
	r.HandleFunc("/booking/{id}", bookingHandler.DeleteBooking).Methods("DELETE")
	r.HandleFunc("/bookings", bookingHandler.GetBookings).Methods("GET")

	return r
}
