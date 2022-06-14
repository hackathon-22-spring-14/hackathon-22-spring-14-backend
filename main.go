package main

import (
	"log"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.Setup(e)

	log.Fatal(e.Start(":3000"))
}
