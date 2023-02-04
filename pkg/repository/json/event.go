package json

import (
	"github.com/growthers/mu-enshita/pkg/types/data"
	"github.com/growthers/mu-enshita/pkg/types/domain"
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
	return domain.Event(e.cache.Event), nil
}
