package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Soldier struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Password     string             `bson:"password" json:"password"`
	Role         string             `bson:"role" json:"role"`
	FortIDs []primitive.ObjectID `bson:"fort_ids" json:"fort_ids"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_token" json:"refresh_token"`
}
