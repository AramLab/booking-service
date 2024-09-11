package routes

import (
	"github.com/AramLab/booking-service/server/booking"
	"github.com/AramLab/booking-service/server/user"
	"github.com/AramLab/booking-service/service"
	"github.com/gorilla/mux"
)

// RegisterRoutes регистрирует маршруты для пользователей и бронирований.
// Он принимает на вход службы UserService и BookingService для обработки запросов,
// затем возвращает HTTP-маршрутизатор с зарегистрированными обработчиками.
func RegisterRoutes(userService service.UserService, bookingService service.BookingService) *mux.Router {
	r := mux.NewRouter()

	// Создаём обработчики для пользователя.
	userHandler := user.NewUserHandlers(userService)

	// Регистрация маршрута для создания нового пользователя.
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")

	// Регистрация маршрута для удаления пользователя по ID.
	r.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Создаём обработчики для бронирования.
	bookingHandler := booking.NewBookingHandlers(bookingService)

	// Регистрация маршрута для создания нового бронирования.
	r.HandleFunc("/booking", bookingHandler.CreateBooking).Methods("POST")

	// Регистрация маршрута для удаления бронирования по ID.
	r.HandleFunc("/booking/{id}", bookingHandler.DeleteBooking).Methods("DELETE")

	// Регистрация маршрута для получения всех бронирований.
	r.HandleFunc("/bookings", bookingHandler.GetBookings).Methods("GET")

	// Возвращаем маршрутизатор с зарегистрированными маршрутами.
	return r
}
