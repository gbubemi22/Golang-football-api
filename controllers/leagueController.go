package controller

import (
	"context"
	"fmt"

	//"log"
	"net/http"
	"time"

	database "github.com/gbubemi22/golang-football-api/initializers"
	"github.com/gbubemi22/golang-football-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var leagueCollection *mongo.Collection = database.OpenCollection(database.Client, "league")
var validate = validator.New()

func Creatleague() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		//defer cancel()
		var league models.League

		if err := c.BindJSON(&league); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//defer cancel()
		validatorErr := validate.Struct(league)
		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr.Error()})
			return

		}

		league.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		league.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		league.ID = primitive.NewObjectID()
		league.League_id = league.ID.Hex()

		result, insertEr := leagueCollection.InsertOne(ctx, league)
		if insertEr != nil {
			msg := fmt.Sprintf(" league was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}

}

func GetLeague() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		leagueId := c.Param("league_id")
		var league models.League

		err := leagueCollection.FindOne(ctx, bson.M{"league_id": leagueId}).Decode(&league)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No league with that id found "})

		}

		c.JSON(http.StatusOK, league)
	}

}
