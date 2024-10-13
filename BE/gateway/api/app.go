package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) InitRoute(g *gin.Engine) {
	InitImageRoute(g)
}

func (a *App) Run(port string) error {

	g := gin.Default()

	g.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	a.InitRoute(g)

	return g.Run(":" + port)
}
