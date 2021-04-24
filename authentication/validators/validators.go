package validators

import (
	"errors"

	"github.com/ygt1qa/microservices/pb"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidUserId      = errors.New("invalid userId")
	ErrEmptyName          = errors.New("name cant be empty")
	ErrEmptyEmail         = errors.New("email cant be empty")
	ErrEmptyPassword      = errors.New("password cant be empty")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

func ValidateSignUp(user *pb.User) error {
	if !bson.IsObjectIdHex(user.Id) {
		return ErrInvalidUserId
	}
	if user.Email == "" {
		return ErrEmptyEmail
	}
	if user.Name == "" {
		return ErrEmptyName
	}
	if user.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}
