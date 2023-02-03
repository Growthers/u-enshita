package domain

import "time"

type Event struct {
	StartAt time.Time
	EndAt   time.Time
}
