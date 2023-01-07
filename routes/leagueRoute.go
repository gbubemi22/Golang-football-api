package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/gbubemi22/golang-football-api/controllers"
) 

func LeagueRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/leagues", controller.CreateLeague())
	incomingRoutes.GET("/leagues/:league_id", controller.GetLeague())
	incomingRoutes.GET("/leagues/", controller.GetLeagues())
	incomingRoutes.PUT("/leagues/:league_id", controller.UpdateFood())
}
       