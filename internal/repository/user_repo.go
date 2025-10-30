package repository

import (
	"context"
	"time"

	"myapp/internal/database"
	"myapp/internal/model"
)

func CreateUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.DB.Collection("users")

	_, err := collection.InsertOne(ctx, user)
	return err
}
