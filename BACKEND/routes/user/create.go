package user

import (
	"goauth/database"
	"goauth/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(ctx echo.Context) error {

	user := entities.User{
		Name:     ctx.FormValue("name"),
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	err := user.PasswordToHash()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Ocoreu um erro ao gerar password")
	}

	db, err := database.DatabaseConnection()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error")
	}

	defer db.Close()

	_, err = user.Create(db)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, user)
}
