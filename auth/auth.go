package auth

import (
	"errors"
	"final/helper"
	"final/model"
	"final/service"
)

var authService = service.AuthService{}

type IAuth interface {
	Login(email, password string) (string, error)
	Register(fullName, email, password string) (string, error)
}

type Auth struct { }

func (Auth) Login(email, password string) (*model.User, error){
	if !helper.IsEmail(email) || !helper.IsRequired(password) {
		return nil, errors.New("Invalid input")
	}
	user, err := authService.GetUserByLogin(email, password)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	return user, nil
}


func (Auth) Register(fullName, email, password string) (string, error){
	if !helper.IsEmail(email) || !helper.IsRequired(password) || !helper.IsRequired(fullName) {
		return "", errors.New("Invalid input")
	}
	return "", nil
}
