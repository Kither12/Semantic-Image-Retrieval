package api

import (
	handler "gateway/handler"

	"github.com/gin-gonic/gin"
)

func InitImageRoute(g *gin.Engine) {
	imageHandler := handler.NewImageHandler()

	sessionRouterGroup := g.Group("api/v1/images")
	sessionRouterGroup.POST("/", imageHandler.Upload)
}
