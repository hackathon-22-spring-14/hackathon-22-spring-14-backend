package router

import (
	"fmt"
	"net/http"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id,omitempty"`
	PassWord string `json:"password,omitempty"`
}

type LoginRequestBody struct {
	ID       string `json:"id,omitempty" form:"id"`
	PassWord string `json:"password,omitempty" form:"password"`
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
	fmt.Println(newUserReq)
	if newUserReq.ID == "" || newUserReq.PassWord == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "項目が空です!")
	}

	//パスワードをハッシュ
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUserReq.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
	}

	newUser := model.User{
		ID:       newUserReq.ID,
		PassWord: string(hashedPass),
	}

	_, message, er := h.r.Signup(newUser)

	if message != "" {
		return c.JSON(http.StatusConflict, message)
	}

	if er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}

	return c.JSON(http.StatusOK, "success creating a user")
}

func (h *userHandler) Login(c echo.Context) error {
	// TODO: implement
	return nil
}
