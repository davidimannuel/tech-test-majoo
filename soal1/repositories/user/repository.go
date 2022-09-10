package user

import (
	"context"
	"time"
)

type UserRepository interface {
	FindOneByUserName(ctx context.Context, username string) (row UserModel, err error)
}

type UserModel struct {
	Id        int
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
