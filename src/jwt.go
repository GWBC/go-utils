package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload[T any] struct {
	jwt.RegisteredClaims
	Data T
}

type Jwt[T any] struct {
	key           []byte
	expiresSecond time.Duration
}

func (j *Jwt[T]) Init(key string, expiresSecond int) {
	j.key = []byte(key)
	j.expiresSecond = time.Duration(expiresSecond) * time.Second
}

func (j *Jwt[T]) Gen(data T) (string, error) {
	payload := JwtPayload[T]{}
	payload.Data = data
	payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.expiresSecond))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)
	ret, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (j *Jwt[T]) Parse(data string) *T {
	token, err := jwt.ParseWithClaims(data, &JwtPayload[T]{}, func(t *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return nil
	}

	if !token.Valid {
		return nil
	}

	payload, ok := token.Claims.(*JwtPayload[T])
	if !ok {
		return nil
	}

	return &payload.Data
}
