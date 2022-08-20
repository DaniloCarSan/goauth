package entities

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) PasswordToHash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err == nil {
		u.Password = string(hashedPassword)
	}

	return err
}

func (u *User) Create(db *sql.DB) (bool, error) {
	stmt, err := db.Prepare("INSERT INTO users(name,email,password) values(?,?,?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return false, err
	}

	code, err := result.LastInsertId()

	if err != nil {
		return true, err
	}

	u.Code = int(code)

	return true, err
}
