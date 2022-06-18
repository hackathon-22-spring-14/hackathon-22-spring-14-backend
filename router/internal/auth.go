package internal

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if sess.Values["userID"] == nil {
			return c.String(http.StatusForbidden, "please login")
		}
		c.Set("userID", sess.Values["userID"].(string))

		return next(c)
	}
}
