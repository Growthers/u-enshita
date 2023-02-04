package json

import (
	"github.com/growthers/mu-enshita/pkg/types/data"
	"github.com/growthers/mu-enshita/pkg/types/domain"
	"time"
)

type EventJSONRepository struct {
	cache      *data.StoreJSON
	readWriter DataReadWriter
}

func NewEventJSONRepository(readWriter DataReadWriter) *EventJSONRepository {
	cache, err := readWriter.read()
	if err != nil {
		panic("failed to read jsonfile")
	}
	return &EventJSONRepository{cache: &cache, readWriter: readWriter}
}

func (e EventJSONRepository) GetEvent() (domain.Event, error) {
	d := domain.Event{}
	ev := e.cache.Event
	d.StartAt, _ = time.Parse(time.RFC3339, ev.StartAt)
	d.EndAt, _ = time.Parse(time.RFC3339, ev.EndAt)
	return d, nil
}
