package view

import (
	"clip/models"
	"fmt"
	"stream/config"
)

type Clip struct {
	Name      string
	Url       string
	Thumbnail string
	Link      string
}

func NewClip(clip models.Clip) Clip {
	clipService := config.Get().ClipService.Address
	fmt.Println(fmt.Sprintf("http://%s/s/%s", clipService, clip.Slug))

	return Clip{
		Name:      clip.Name,
		Url:       clip.Url,
		Thumbnail: fmt.Sprintf("/static/images/%s.jpg", clip.Slug),
		Link:      fmt.Sprintf("%s/s/%s", clipService, clip.Slug),
	}
}

func NewClips(clips []models.Clip) (clipViews []Clip) {
	clipViews = make([]Clip, len(clips))
	for i, c := range clips {
		clipViews[i] = NewClip(c)
	}

	return
}
