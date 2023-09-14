package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=2,max=50"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	// Role      *string            `json:"role" validate:"eq=MENTOR|eq=INTERN|eq="`
}
