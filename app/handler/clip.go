package handler

import (
	"fmt"
	"stream/app/entity"
	"stream/app/view"
	"stream/config"

	"github.com/gin-gonic/gin"
)

type streamHandler struct {
	clip entity.Clip
}

type Presenter struct {
	Clips              []view.Clip
	ClipServiceAddress string
}

func (h streamHandler) Index(c *gin.Context) {
	presenter := Presenter{}

	clips, err := h.clip.GetLatestClips()
	if err != nil {
		fmt.Println(err)
		c.HTML(200, "index.html", presenter)
		return
	}

	presenter.Clips = view.NewClips(clips)
	presenter.ClipServiceAddress = config.Get().ClipService.Address
	c.HTML(200, "index.html", presenter)
}
