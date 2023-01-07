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

	//"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var teamCollection *mongo.Collection = database.OpenCollection(database.Client, "team")


// var validate = validator.New()

func GetTeams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := teamCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err!= nil {
           c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fatchinf teams"})
	}
	var allTeams []bson.M
	if err = result.All(ctx, &allTeams); err !=nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, allTeams)
     return

}

}


func GetTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		teamId := c.Param("food_id")
		var team models.Team

		err := teamCollection.FindOne(ctx, bson.M{"team_id": teamId}).Decode(&team)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, team)
	}
}


func CreateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var team models.Team
		var league models.League

		if err := c.BindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(team)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := leagueCollection.FindOne(ctx, bson.M{"league_id": team.League_id}).Decode(&league)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("league was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		team.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		team.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		team.ID = primitive.NewObjectID()
		team.Team_id = team.ID.Hex()
		

		result, insertErr := teamCollection.InsertOne(ctx, team)
		if insertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}


func UpdateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var league models.League
		var team models.Team

		teamId := c.Param("team_id")

		if err := c.BindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if team.Team != nil {
			updateObj = append(updateObj, bson.E{"team", team.Team})
		}

		if team.NickName != nil {
			updateObj = append(updateObj, bson.E{"nickName", team.NickName})
		}

		if team.Team_image != nil {
			updateObj = append(updateObj, bson.E{"food_image", team.Team_image})
		}

		if team.League_id != nil {
			err := leagueCollection.FindOne(ctx, bson.M{"league_id": team.League_id}).Decode(&league)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			// updateObj = append(updateObj, bson.E{"league", food.Price})
		}

		team.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", team.Updated_at})

		upsert := true
		filter := bson.M{"food_id": teamId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := teamCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprint("foot item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}