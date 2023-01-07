package controller

import (
	"context"
	"fmt"
	"log"

	//"log"
	"net/http"
	"time"

	database "github.com/gbubemi22/golang-football-api/database"
	"github.com/gbubemi22/golang-football-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)


var leagueCollection *mongo.Collection = database.OpenCollection(database.Client, "league")

var validate = validator.New()

func GetLeagues() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := leagueCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing league items"})
		}
		var allLeagues []bson.M
		if err = result.All(ctx, &allLeagues); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allLeagues)
		fmt.Println(allLeagues)
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the league"})
		}
		c.JSON(http.StatusOK, league)
		fmt.Println(league)
	}
}


func CreateLeague() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var league models.League

		if err := c.BindJSON(&league); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(league)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		

		league.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		league.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		league.ID = primitive.NewObjectID()
		league.League_id = league.ID.Hex()

		result, insertErr := leagueCollection.InsertOne(ctx, league)

		if insertErr != nil {
			msg := fmt.Sprintf("league item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}





func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		var league models.League

		leagueId := c.Param("league_id")

		if err := c.BindJSON(&league); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
 

		
		

		var updateObj primitive.D

		if league.League_name != nil {
			updateObj = append(updateObj, bson.E{"league_name", league.League_name})
		}

		if league.Location != nil {
			updateObj = append(updateObj, bson.E{"location", league.Location})
		}

		if league.League_image != nil {
			updateObj = append(updateObj, bson.E{"league_image", league.League_image})
		}

		

		

		league.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", league.Updated_at})

		upsert := true
		filter := bson.M{"league_id": leagueId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := leagueCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprint("league item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}





