package main

import (
	"final/model"
	"final/service"
	"fmt"
)

func main(){
	authService := service.AuthService{}

	user := &model.User{
		FullName: "Margot Robbie",
		Email: "m.robbie@astanait.edu.kz",
		Password: "mro",
	}

	err := authService.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User was added")
	}
}
