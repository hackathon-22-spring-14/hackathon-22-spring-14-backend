package router

import (
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository/impl"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, db *sqlx.DB) {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sh := NewStampHandler(impl.NewStampRepository(db)) //いくら用メモ---shには、dbの入ったstampRepositoryの入ったstampHandlerが入っているけど、StampHandlerで返り値を指定しているから、GetStampsを呼び出せる。

	api := e.Group("/api")
	apiStamps := api.Group("/stamps")

	apiStamps.GET("", sh.GetStamps)
	apiStamps.POST("", sh.PostStamp)
	apiStamps.GET("/:stampID", sh.GetStamp)
	apiStamps.DELETE("/:stampID", sh.DeleteStamp)
}
