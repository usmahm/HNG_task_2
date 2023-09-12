package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName *string            `json:"first_name" validate:"required,min=2,max=50"`
	LastName  *string            `json:"last_name" validate:"required,min=2,max=50"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Role      *string            `json:"role" validate:"eq=MENTOR|eq=INTERN|eq="`
}
