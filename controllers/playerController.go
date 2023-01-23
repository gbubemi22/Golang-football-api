package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gbubemi22/golang-football-api/database"
	"github.com/gbubemi22/golang-football-api/models"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var validate = validator.New()

var playerCollection *mongo.Collection = database.OpenCollection(database.Client, "player")

func GetPlayers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := playerCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fatchinf teams"})
		}
		var allPlayers []bson.M
		if err = result.All(ctx, &allPlayers); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allPlayers)
		return
	}
}

func GetPlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		playerId := c.Param("player_id")
		var player models.Player

		err := playerCollection.FindOne(ctx, bson.M{"player_id": playerId}).Decode(&player)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the player"})
		}
		c.JSON(http.StatusOK, player)
	}
}

func CreatePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var player models.Player
		var team models.Team

		if err := c.BindJSON(&player); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(player)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := teamCollection.FindOne(ctx, bson.M{"team_id": player.Team_id}).Decode(&team)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("team was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		player.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		player.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		player.ID = primitive.NewObjectID()
		player.Player_id = player.ID.Hex()

		result, insertErr := playerCollection.InsertOne(ctx, player)
		if insertErr != nil {
			msg := fmt.Sprintf("Player was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdatePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var team models.Team
		var player models.Player

		playerId := c.Param("player_id")

		if err := c.BindJSON(&player); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if player.Player_name != nil {
			updateObj = append(updateObj, bson.E{"player_name", player.Player_name})
		}

		if player.Position != nil {
			updateObj = append(updateObj, bson.E{"position", player.Position})
		}

		if player.Nationality != nil {
			updateObj = append(updateObj, bson.E{"nationality", player.Nationality})
		}

		if player.Number != nil {
			updateObj = append(updateObj, bson.E{"number", player.Number})
		}

		
		if player.Image  != nil {
			updateObj = append(updateObj, bson.E{"image ", player.Image })
		}

		if player.Team_id  != nil {
			err := teamCollection.FindOne(ctx, bson.M{"team_id": player.Team_id}).Decode(&team)
			

               defer cancel()


			if err != nil {
				msg := fmt.Sprintf("team was not found")
                    c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
                    return
               }

			player.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", player.Updated_at})

			 upsert := true
			 filter := bson.M{"player_id": playerId}

			 opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := playerCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)
	
			if err != nil {
				msg := fmt.Sprint("player  update failed")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			c.JSON(http.StatusOK, result)

			
		}


	}

}
