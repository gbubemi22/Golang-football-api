package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/gbubemi22/golang-football-api/controllers"
) 

func LeagueRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/api/leagues", controller.Creatleague())
	incomingRoutes.GET("/api/leagues/league_id", controller.GetLeague())
}