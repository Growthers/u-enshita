package domain

import "time"

type Event struct {
	StartAt time.Time
	EndAt   time.Time
}

func NewEvent(startAt time.Time, endAt time.Time) *Event {
	return &Event{StartAt: startAt, EndAt: endAt}
}
