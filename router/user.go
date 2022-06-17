package router

import (
	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	PassWord string    `db:"password"`
}

type UserHandler interface {
	// POST /users/signup
	Signup(c echo.Context) error
	// POST /users/login
	Login(c echo.Context) error
}

type userHandler struct {
	r repository.UserRepository
}

func NewUserHandler(r repository.UserRepository) UserHandler {
	return &userHandler{r}
}

func (h *userHandler) Signup(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *userHandler) Login(c echo.Context) error {
	// TODO: implement
	return nil
}
