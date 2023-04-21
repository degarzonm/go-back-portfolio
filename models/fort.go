package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fort struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FortID    *string            `bson:"fort_id" json:"fort_id"`
	Name      *string            `bson:"name" json:"name"`
	Location  *string            `bson:"location" json:"location"`
	Commander *string            `bson:"commander," json:"commander"`
	AgendaID  *string            `bson:"agenda_id," json:"agenda_id"`
	PlanID    *string            `bson:"plan_id," json:"plan_id"`
}
