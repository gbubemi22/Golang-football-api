package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type League struct {
	ID           primitive.ObjectID `bson:"_id"`
	League_name  *string            `json:"league_name" validate:"required"`
	Location     *string            `json:"location" validate:"required`
	League_image *string            `json:"league_image" validate:"required`
	League_id    string             `json:"league_id"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}
