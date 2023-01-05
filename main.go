package main

import (
	"github.com/gbubemi22/golang-football-api/initializers"
	"github.com/gin-gonic/gin"
	routes "github.com/gbubemi22/golang-football-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}





func main() {


	router := gin.New()
router.Use(gin.Logger())
//routes.UserRoutes(router)

routes.LeagueRoutes(router)
	// r := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this API IS LIVE",
		})
	})
	
	 
	router.Run() // listen and serve on 0.0.0.0:8080
}
