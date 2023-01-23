package main

import (
	"os"

	"github.com/gbubemi22/golang-football-api/database"
	routes "github.com/gbubemi22/golang-football-api/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var leagueCollection *mongo.Collection = database.OpenCollection(database.Client, "league")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5003"
	}

	router := gin.New()
	router.Use(gin.Logger())
	//routes.UserRoutes(router)

	routes.LeagueRoutes(router)
	routes.TeamRoutes(router)
	routes.PlayerRoutes(router)


	router.Run(":" + port)
}
