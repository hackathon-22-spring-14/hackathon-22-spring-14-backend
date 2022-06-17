package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hackathon-22-spring-14/hackathon-22-spring-14-backend/router"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sqlx.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := db.DB.Ping(); err == nil {
			break
		} else if i == 9 {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(i+1))
	}

    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	e := echo.New()
	router.Setup(e, db, cfg)

	e.Logger.Fatal(e.Start(":3000"))
}

func getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
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
