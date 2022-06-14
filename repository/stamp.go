package repository

import "github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"

type StampRepository interface {
	FindAll() ([]model.Stamp, error)
}
