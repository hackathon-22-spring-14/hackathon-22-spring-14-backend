package impl

import (
	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	PassWord string    `db:"password"`
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Signup() error {
	return nil
}

func (r *userRepository) Login() error {
	return nil
}
