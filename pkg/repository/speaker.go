package repository

import "github.com/growthers/mu-enshita/pkg/types/domain"

type SpeakerRepository interface {
	FindAllSpeakers() ([]domain.Speaker, error)
	FindSpeakersByID(id string) (domain.Speaker, error)
	UpdateSpeaker(id string, speaker domain.Speaker) (domain.Speaker, error)
}
