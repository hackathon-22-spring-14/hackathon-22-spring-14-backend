package repository

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
)

type UserRepository interface {
	Signup(newUser model.User) (model.User, string, error)
	Login() error
}
