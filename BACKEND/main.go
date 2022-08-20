package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {

		ctx.JSON(http.StatusOK, "")
	})

	add := fmt.Sprintf("%s%s", os.Getenv("HOST_NAME"), os.Getenv("HOST_PORT"))
	e.Logger.Fatal(e.Start(add))
}
