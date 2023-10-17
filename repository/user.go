package repository

import (
	"context"
	"rest-ws/models"
)

type UserResitory interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
}

var implementation UserResitory

func SetRepository(repository UserResitory) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}