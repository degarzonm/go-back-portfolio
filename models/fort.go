package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fort struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Location   string             `bson:"location" json:"location"`
	Commander  primitive.ObjectID `bson:"commander,omitempty" json:"commander"`
	AgendaID   primitive.ObjectID `bson:"agenda_id,omitempty" json:"agenda_id"`
	PlanID     primitive.ObjectID `bson:"plan_id,omitempty" json:"plan_id"`
}

