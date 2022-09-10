package middlewares

import (
	"context"
	"errors"
	"fmt"
	"majoo/soal1/configs"
	"majoo/soal1/controllers"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserID   int
	Username string
	jwt.StandardClaims
}

type JWT struct {
	Config configs.JWT
}

func SetupJWT(conf configs.JWT) *JWT {
	return &JWT{Config: conf}
}

func (j JWT) getSecret() []byte {
	return []byte(j.Config.Secret)
}

func (j JWT) GenerateToken(ctx context.Context, userID int, username string) (res string, err error) {
	claims := &JwtClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.Config.Expiration)).Unix(), // miliseconds
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err = token.SignedString(j.getSecret())

	return
}

func (j JWT) Validate(ctx context.Context, tokenStr string) (claims JwtClaims, err error) {
	_, err = jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return j.getSecret(), nil
	})
	if err != nil {
		return
	}

	return
}

func (j JWT) getHTTPBearerToken(h http.Header) (string, error) {
	authH := h.Get("Authorization")
	if !strings.Contains(authH, "Bearer") {
		return "", errors.New("invalid token")
	}
	return strings.Replace(authH, "Bearer ", "", -1), nil
}

func (j JWT) MuxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token, err := j.getHTTPBearerToken(r.Header)
		if err != nil {
			controllers.WriteResponse(w, http.StatusUnauthorized, err.Error(), nil, nil)
			return
		}

		jwtClaims, err := j.Validate(ctx, token)
		if err != nil {
			controllers.WriteResponse(w, http.StatusUnauthorized, err.Error(), nil, nil)
			return
		}

		ctx = context.WithValue(ctx, CtxKeyUserID, jwtClaims.UserID)
		ctx = context.WithValue(ctx, CtxKeyUserName, jwtClaims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
