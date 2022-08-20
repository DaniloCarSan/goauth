package user

import (
	"goauth/database"
	"goauth/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AllUsers(ctx echo.Context) error {

	db, err := database.DatabaseConnection()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error ao se conectar a base de dados")
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var user entities.User

	var users []entities.User

	for rows.Next() {
		rows.Scan(&user.Code, &user.Name, &user.Email)
		users = append(users, user)
	}

	return ctx.JSON(http.StatusOK, users)
}
