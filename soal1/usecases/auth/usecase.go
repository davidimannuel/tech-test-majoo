package auth

import "context"

type AuthUsecase interface {
	Login(ctx context.Context, ent AuthEntity) (token string, err error)
}

type AuthEntity struct {
	UserName string
	Password string
}
