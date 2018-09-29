package entity

import (
	"clip/models"
	"clip/repo"
	"clip/utilities/uer"
)

type clipEntity struct {
	clipRepo repo.IClip
}

type Clip interface {
	GetLatestClips() (clips []models.Clip, err error)
}

func NewClip(clipRepo repo.IClip) Clip {
	return &clipEntity{
		clipRepo: clipRepo,
	}
}

func (c clipEntity) GetLatestClips() (clips []models.Clip, err error) {
	offset := 0
	limit := 4
	clips, err = c.clipRepo.GetLatest(offset, limit)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}
