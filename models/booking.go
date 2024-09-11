package models

import "time"

type Booking struct {
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Start_time time.Time `json:"start_time"`
	End_time   time.Time `json:"end_time"`
}
