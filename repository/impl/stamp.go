package impl

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/jmoiron/sqlx"
)

type stampRepository struct {
	db *sqlx.DB
}

func NewStampRepository(db *sqlx.DB) repository.StampRepository {
	return &stampRepository{db}
}

func (r *stampRepository) FindAll() ([]model.Stamp, error) {
	return []model.Stamp{}, nil
}
