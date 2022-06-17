package router

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	PassWord string    `json:"password,omitempty"`
}

type LoginRequestBody struct {
	Name     string `json:"name,omitempty"`
	PassWord string `json:"password,omitempty"`
}

type LoginResponseBody struct {
	ID uuid.UUID `json:"id,omitempty"`
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
	newUserReq := LoginRequestBody{}
	if er := c.Bind(&newUserReq); er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}

	if newUserReq.Name == "" || newUserReq.PassWord == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "項目が空です!")
	}

	//パスワードをハッシュ
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUserReq.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	//IDを生成
	newUserID := uuid.New()

	newUser := model.User{
		ID:       newUserID,
		Name:     newUserReq.Name,
		PassWord: string(hashedPass),
	}

	addedUser, message, er := h.r.Signup(newUser)

	if message != "" {
		return c.JSON(http.StatusConflict, message)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}

	res := LoginResponseBody{ID: addedUser.ID}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandler) Login(c echo.Context) error {
	// TODO: implement
	return nil
}
