package json

import (
	"encoding/json"
	"errors"
	"github.com/growthers/mu-enshita/pkg/types/data"
	"github.com/growthers/mu-enshita/pkg/types/domain"
)

type SpeakerJSONRepository struct {
	cache      *data.StoreJSON
	readWriter DataReadWriter
}

func NewSpeakerJSONRepository(readWriter DataReadWriter) *SpeakerJSONRepository {
	cache, err := readWriter.read()
	if err != nil {
		panic("failed to read jsonfile")
	}

	return &SpeakerJSONRepository{cache: &cache, readWriter: readWriter}
}

func (r *SpeakerJSONRepository) writeData(d data.StoreJSON) error {
	bytes, e := json.Marshal(d)
	if e != nil {
		return e
	}

	err := r.readWriter.write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (r *SpeakerJSONRepository) typeConverter(d data.SpeakerStoreJSON) domain.Speaker {
	res := domain.Speaker(d)
	return res
}

func (r *SpeakerJSONRepository) FindAllSpeakers() ([]domain.Speaker, error) {
	res := make([]domain.Speaker, len(r.cache.Speakers))

	for i, v := range r.cache.Speakers {
		res[i] = r.typeConverter(v)
	}

	return res, nil
}

func (r *SpeakerJSONRepository) FindSpeakersByID(id string) (domain.Speaker, error) {
	for _, v := range r.cache.Speakers {
		if v.ID == id {
			return r.typeConverter(v), nil
		}
	}

	return domain.Speaker{}, errors.New("NotFound")
}

func (r *SpeakerJSONRepository) UpdateSpeaker(id string, speaker domain.Speaker) (domain.Speaker, error) {
	// キャッシュから更新すべきデータを読み出す
	d := r.cache.Speakers
	for i, v := range r.cache.Speakers {
		if v.ID == id {
			d[i] = data.SpeakerStoreJSON(speaker)
		}
	}

	// キャッシュを更新
	r.cache.Speakers = d

	// ファイルに現在のキャッシュの内容を書き込み
	err := r.writeData(*r.cache)
	if err != nil {
		return domain.Speaker{}, err
	}

	return speaker, nil
}
