package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/model"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
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

type Me struct {
	UserID string `json:"user_id,omitempty"  db:"user_id"`
}

type UserHandler interface {
	// POST /users/signup
	Signup(c echo.Context) error
	// POST /users/login
	Login(c echo.Context) error
	// GET /users/whoami
	GetWhoAmIHandler(c echo.Context) error
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

	sess, err := session.Get("sessions", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userID"] = newUserReq.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusCreated, "success creating a user")
}

func (h *userHandler) Login(c echo.Context) error {
	loginReq := LoginRequestBody{}
	if er := c.Bind(&loginReq); er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}
	if loginReq.ID == "" || loginReq.PassWord == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "項目が空です!")
	}

	user := model.User{
		ID:       loginReq.ID,
		PassWord: loginReq.PassWord,
	}

	message, er := h.r.Login(user)
	if er != nil {
		fmt.Println(er)
		if message != "" {
			return c.JSON(http.StatusForbidden, message)
		}
		if errors.Is(er, bcrypt.ErrMismatchedHashAndPassword) {
			return c.JSON(http.StatusForbidden, "the password is wrong")
		}
		
		return c.JSON(http.StatusInternalServerError, er)
		
	}
	// TODO: implement
	return c.JSON(http.StatusOK, "success loging in")
}

func (h *userHandler) GetWhoAmIHandler(c echo.Context) error {
	sess, _ := session.Get("sessions", c)
  
	return c.JSON(http.StatusOK, Me{
		UserID: sess.Values["userID"].(string),
	})
}
