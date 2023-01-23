package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Player struct {
	ID          primitive.ObjectID `bson:"id"`
	Player_name *string            `json:"player_name" validate:"required"`
	Position    *string            `json:"position" validate:"required"`
	Nationality *string            `json:"nationality" validate:"required"`
	Number      *string            `json:"number" validate:"required"`
	IsCaptain   bool               `json:"isCaptain" validate:"default:false"`
	Image       *string            `json:"image" `
	Player_id   string            `json:"player_id"`
	Team_id     *string            `json:"team_id" validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
