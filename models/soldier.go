package models

import (
	"github.com/degarzonm/go-back-portfolio/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Soldier struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       *string            `bson:"user_id" json:"user_id"`
	Name         *string            `bson:"name" json:"name"`
	City         *string            `bson:"city" json:"city"`
	Password     *string            `bson:"password" json:"password"`
	Role         middleware.Role    `bson:"role" json:"role"`
	FortIDs      []string           `bson:"fort_ids" json:"fort_ids"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_token" json:"refresh_token"`
}

type SoldierLoginResponse struct {
	UserID       *string  `json:"user_id"`
	Name         *string  `json:"name"`
	City         *string  `json:"city"`
	FortIDs      []string `json:"fort_ids"`
	Token        *string  `json:"token"`
	RefreshToken *string  `json:"refresh_token"`
}
