package repository

import (
	"Arc/model"
	"context"
)

type postgress interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserID(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

