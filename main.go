package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/router"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := sqlx.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err)
	}

	router.Setup(e, db)

	log.Fatal(e.Start(":3000"))
}

func getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		getEnvOrDefault("MYSQL_USER", "root"),
		getEnvOrDefault("MYSQL_PASSWORD", "password"),
		getEnvOrDefault("MYSQL_HOST", "localhost"),
		getEnvOrDefault("MYSQL_PORT", "3306"),
		getEnvOrDefault("MYSQL_DATABASE", "hackathon"), // TODO: サービス名に変える
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
