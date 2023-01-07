package routes

import (
	controller "github.com/gbubemi22/golang-football-api/controllers"
	"github.com/gin-gonic/gin"
)

func TeamRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/teams/:team_id",controller.GetTeam())
     incomingRoutes.GET("/teams", controller.GetTeams())
     incomingRoutes.POST("/teams", controller.CreateTeam())
    incomingRoutes.PUT("/teams/:team_id/", controller.UpdateTeam())
}















