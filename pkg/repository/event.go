package repository

import "github.com/growthers/mu-enshita/pkg/types/domain"

type EventRepository interface {
	GetEvent() (domain.Event, error)
}
