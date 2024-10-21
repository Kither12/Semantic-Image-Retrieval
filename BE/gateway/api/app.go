package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

	gin.SetMode(viper.GetString("MODE"))
	g := gin.Default()

	g.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
	)

	a.InitRoute(g)

	return g.Run(":" + port)
}
