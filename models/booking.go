package models

import "time"

// Booking представляет информацию о бронировании.
type Booking struct {
	ID         int       `json:"id"`
	User_id    int       `json:"user_id" validate:"required"`
	Start_time time.Time `json:"start_time" validate:"required"`
	End_time   time.Time `json:"end_time" validate:"required,gtfield=Start_time"`
}
