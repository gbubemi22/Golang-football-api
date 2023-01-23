package routes

import (
	controller "github.com/gbubemi22/golang-football-api/controllers"

	"github.com/gin-gonic/gin"
)


func PlayerRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/players", controller.GetPlayers())

		incomingRoutes.GET("/players/:player_id", controller.GetPlayer())

	incomingRoutes.POST("players", controller.CreatePlayer())
	incomingRoutes.PUT("players/:player_id", controller.UpdatePlayer())

}