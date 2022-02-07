package entity

import "time"

type Mail struct {
	Id          string
	Destination string
	Description string
	Message     string
	Date        time.Time
}
