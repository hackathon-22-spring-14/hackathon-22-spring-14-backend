package router

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/repository/impl"
	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/router/internal"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/srinathgs/mysqlstore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func Setup(e *echo.Echo, db *sqlx.DB, cfg aws.Config) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))

	sh := NewStampHandler(impl.NewStampRepository(db, impl.NewStampStrage(cfg))) //いくら用メモ---shには、dbの入ったstampRepositoryの入ったstampHandlerが入っているけど、StampHandlerで返り値を指定しているから、GetStampsを呼び出せる。
	uh := NewUserHandler(impl.NewUserRepository(db))                             //いくら用メモ---uhには、dbの入ったuserRepositoryの入ったuserHandlerが入っているけど、UserHandlerで返り値を指定しているから、Signupとかを呼び出せる。

	api := e.Group("/api")
	
	apiStamps := api.Group("/stamps")
	apiStamps.Use(internal.CheckLogin)
	apiStamps.GET("", sh.GetStamps)
	apiStamps.POST("", sh.PostStamp)
	apiStamps.GET("/:stampID", sh.GetStamp)
	apiStamps.DELETE("/:stampID", sh.DeleteStamp)

	apiUsers := api.Group("/users")
	apiUsers.POST("/signup", uh.Signup)
	apiUsers.POST("/login", uh.Login)
	apiUsers.GET("/whoami", uh.GetWhoAmIHandler, internal.CheckLogin)
}
