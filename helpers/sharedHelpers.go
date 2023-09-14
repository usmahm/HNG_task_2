package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsObjectID(str string) bool {
	_, err := primitive.ObjectIDFromHex(str)
	return err == nil
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	return ctx, cancel
}
