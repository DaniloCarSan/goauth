package main

import (
	"fmt"
	"log"
	"os"

	"goauth/routes/user"

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

	e.GET("/users", user.AllUsers)
	e.POST("/users", user.Create)

	add := fmt.Sprintf("%s%s", os.Getenv("HOST_NAME"), os.Getenv("HOST_PORT"))
	e.Logger.Fatal(e.Start(add))
}
