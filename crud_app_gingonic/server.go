package main

import (
	"projects/crud_app_gingonic/handler"
	"projects/crud_app_gingonic/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService    = service.New()
	videoController handler.VideoController = handler.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
