package router

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository/impl"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, db *sqlx.DB) {
	sh := NewstampHandler(impl.NewStampRepository(db))

	api := e.Group("/api")
	apiStamps := api.Group("/stamps")

	apiStamps.GET("", sh.GetStamps)
}
