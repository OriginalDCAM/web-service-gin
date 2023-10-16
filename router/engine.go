package router

import (
	"fmt"

	"github.com/OrignalDCAM/web-service-gin/database"
	handler "github.com/OrignalDCAM/web-service-gin/handlers"
	"github.com/gin-gonic/gin"
)

func Run(Port string) {
	database.InitDb()

	engine := gin.Default()

	engine.GET("/games", func(c *gin.Context) {
		handler.GetAllGames(c, database.GetDB())
	})

	engine.GET("/game/:id", func(c *gin.Context) {
		handler.GetGameByID(c, database.GetDB())
	})

	engine.POST("/games", func(c *gin.Context) {
		handler.CreateNewGame(c, database.GetDB())
	})

	engine.Use(gin.Recovery())

	fmt.Println(engine.Run(fmt.Sprintf(":%s", Port)))
}