package json

import (
	"github.com/growthers/mu-enshita/lib/inmemory"
	"github.com/growthers/mu-enshita/pkg/types/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var JST, _ = time.LoadLocation("Asia/Tokyo")

const SAMPLE_JSON_DATA = `{"event":{"startAt":"2023-01-31T10:10:10+09:00","endAt":"2023-01-31T10:50:00+09:00"},"speakers":[{"id":"123-456668-12121","name":"あいうえお","title":"デボビゲゴ","social":[{"serviceName":"GitHub","accountName":"@lame"}],"duration":15,"breakTime":false,"order":3,"cancel":true}]}`

func TestGetEvent(t *testing.T) {
	dummyFile := inmemory.DummyFile{}

	dummyFile.Write([]byte(SAMPLE_JSON_DATA))
	dummyFile.Seek(0, 0)

	repo := NewEventJSONRepository(DataReadWriter{&dummyFile})
	start, _ := time.Parse(time.RFC3339, "2023-01-31T10:10:10+09:00")
	end, _ := time.Parse(time.RFC3339, "2023-01-31T10:50:00+09:00")
	expected := domain.Event{
		StartAt: start,
		EndAt:   end,
	}

	actual, _ := repo.GetEvent()

	assert.Equal(t, expected, actual)
}
