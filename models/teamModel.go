package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID         primitive.ObjectID `bson:"id"`
	Team       *string            `json:"team" validate:"required`
	NickName   *string            `json:"nickName" validate:"required"`
	Team_id    string             `json:"team_id"`
	Team_image *string            `json:"team_image" validate:"required"`
	League_id  *string            `json:"league_id" validate:"required`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
