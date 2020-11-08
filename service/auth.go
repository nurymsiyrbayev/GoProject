package service

import (
	_ "context"
	"errors"
	_ "errors"
	"final/database"
	"final/model"
)

type AuthService struct { }

func (AuthService) GetUserByLogin(email, password string) (*model.User, error) {
	conn := database.GetInstance()

	user := model.User{}
	err := conn.Get(&user, "SELECT id, full_name, email, is_admin FROM users WHERE email=$1 AND password=$2", email, password)

	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}

func (AuthService) RegisterUser(user *model.User) error {
	conn := database.GetInstance()

	_, err := conn.NamedExec("INSERT INTO users (full_name, email, password) VALUES (:full_name, :email, :password)", &user)
	if err != nil {
		return errors.New("Registration failed")
	}

	return nil
}
