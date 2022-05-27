package handler

import (
	"projects/crud_app_gingonic/model"
	"projects/crud_app_gingonic/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) model.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) model.Video {
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
