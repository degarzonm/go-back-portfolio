package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agenda struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Fort string             `bson:"fort_id" json:"fort_id"`
	File string             `bson:"file" json:"file"`
}
