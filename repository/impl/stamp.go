package impl

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
)

type stampRepository struct {
}

func NewStampRepository() repository.StampRepository {
	return &stampRepository{}
}

func (r *stampRepository) FindAll() ([]model.Stamp, error) {
	return []model.Stamp{}, nil
}
