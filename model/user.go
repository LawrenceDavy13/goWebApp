package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-redis/redis"
)
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid login")
)

func AuthenticateUser(username, password string) error {
	hash, err := client.Get("user:" + username).Bytes()
	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return ErrInvalidLogin
	}
	return nil
}

func RegisterUser(username, passoword string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return client.Set("user:"+username, hash, 0).Err()
}