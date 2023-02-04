package json

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/growthers/mu-enshita/pkg/lib/inmemory"
	"github.com/growthers/mu-enshita/pkg/types/domain"
)

func TestFindAllSpeakers(t *testing.T) {
	dummyFile := inmemory.DummyFile{}
	_, err := dummyFile.Write([]byte(SAMPLE_JSON_DATA))
	if err != nil {
		return
	}
	_, err = dummyFile.Seek(0, 0)
	if err != nil {
		return
	}

	repo := NewSpeakerJSONRepository(DataReadWriter{&dummyFile})
	expect := []domain.Speaker{
		{
			ID:    "123-456668-12121",
			Name:  "あいうえお",
			Title: "デボビゲゴ",
			Social: []struct {
				ServiceName string
				AccountName string
			}{{ServiceName: "GitHub", AccountName: "@lame"}},
			Duration:  15,
			BreakTime: false,
			Order:     3,
			Cancel:    true,
		},
	}

	all, err := repo.FindAllSpeakers()
	if err != nil {
		return
	}

	assert.Equal(t, expect, all)
}

func TestFindSpeakersByID(t *testing.T) {
	dummyFile := inmemory.DummyFile{}
	_, err := dummyFile.Write([]byte(SAMPLE_JSON_DATA))
	if err != nil {
		return
	}
	_, err = dummyFile.Seek(0, 0)
	if err != nil {
		return
	}

	repo := NewSpeakerJSONRepository(DataReadWriter{&dummyFile})
	expect := domain.Speaker{
		ID:    "123-456668-12121",
		Name:  "あいうえお",
		Title: "デボビゲゴ",
		Social: []struct {
			ServiceName string
			AccountName string
		}{{ServiceName: "GitHub", AccountName: "@lame"}},
		Duration:  15,
		BreakTime: false,
		Order:     3,
		Cancel:    true,
	}

	s, err := repo.FindSpeakersByID("123-456668-12121")
	if err != nil {
		return
	}

	assert.Equal(t, expect, s)
}

func TestUpdateSpeaker(t *testing.T) {
	dummyFile := inmemory.DummyFile{}
	_, err := dummyFile.Write([]byte(SAMPLE_JSON_DATA))
	if err != nil {
		return
	}
	_, err = dummyFile.Seek(0, 0)
	if err != nil {
		return
	}

	repo := NewSpeakerJSONRepository(DataReadWriter{&dummyFile})
	expect := domain.Speaker{
		ID:    "123-456668-12121",
		Name:  "あいうえお",
		Title: "タイトル",
		Social: []struct {
			ServiceName string
			AccountName string
		}{{ServiceName: "GitHub", AccountName: "@lame"}},
		Duration:  10,
		BreakTime: false,
		Order:     3,
		Cancel:    false,
	}

	s, err := repo.UpdateSpeaker("123-456668-12121", expect)
	if err != nil {
		return
	}

	assert.Equal(t, expect, s)
}
