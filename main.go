package main

import (
	engine "github.com/OrignalDCAM/web-service-gin/router"
	_ "github.com/lib/pq"
)




func main() {
	engine.Run("8001")

	// router := gin.Default()
	// router.GET("/games", getAllGames)
	// router.GET("/game/:id", getGameByID)

	// router.Run("localhost:8001")
}
