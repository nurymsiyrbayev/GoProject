package auth

// Facade

import (
	"errors"
	"final/auth/state"
	"final/helper"
	"final/model"
	"final/service"
)

var authService = &service.AuthService{}
var Session *state.Context = &state.Context{}

type IAuth interface {
	Login(email, password string) (*model.User, error)
	Register(fullName, email, password string) error
	Logout()
}

type Auth struct { }

func (Auth) Login(email, password string) (*model.User, error){
	// validate input
	if !helper.IsEmail(email) || !helper.IsRequired(password) {
		return nil, errors.New("Invalid input")
	}

	// get user be credentials
	user, err := authService.GetUserByLogin(email, password)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	// change state to authorized
	authorizedState := &state.AuthorizedState{ User: user}
	authorizedState.Execute(Session)

	return user, nil
}

func (Auth) Register(fullName, email, password string) error {
	// validate input
	if !helper.IsEmail(email) || !helper.IsRequired(password) || !helper.IsRequired(fullName) {
		return errors.New("Invalid input")
	}

	// create new user
	err := authService.RegisterUser(&model.User{
		FullName: fullName,
		Email: email,
		Password: password,
	})

	if err != nil {
		return errors.New("Something went wrong")
	}

	return nil
}


func (Auth) Logout(){
	// change state to unauthorized
	unauthorizedState := &state.UnauthorizedState{}
	unauthorizedState.Execute(Session)
}
