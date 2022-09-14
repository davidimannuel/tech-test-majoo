package auth

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"majoo/soal1/controllers/middlewares"
	"majoo/soal1/repositories/user"
)

type authUsecase struct {
	jwt      *middlewares.JWT
	userRepo user.UserRepository
}

func NewAuthUsecase(jwt *middlewares.JWT, userRepository user.UserRepository) AuthUsecase {
	return &authUsecase{
		jwt:      jwt,
		userRepo: userRepository,
	}
}

func (uc *authUsecase) Login(ctx context.Context, ent AuthEntity) (token string, err error) {
	user, err := uc.userRepo.FindOneByUserName(ctx, ent.UserName)
	if err != nil {
		return
	}

	encryptedPassword := fmt.Sprintf("%x", md5.Sum([]byte(ent.Password)))
	if user.Password != string(encryptedPassword) {
		err = errors.New("invalid password")
		if err != nil {
			return
		}
	}

	return uc.jwt.GenerateToken(ctx, user.Id, user.UserName)
}
