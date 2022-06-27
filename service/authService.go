package service

import (
	"backend/database"
	"backend/model"
	"backend/repository"
	"backend/util"
	"fmt"
)

func SignUp(user model.User) (string, *util.RestError) {
	if len(user.Username) < 4 || len(user.Username) > 16 {
		return "", util.NewRestErrBadRequest("Username needs to have between 4 and 16 characters")
	}
	if len(user.Password) < 6 {
		return "", util.NewRestErrBadRequest("Password must be atleast 6 characters long")
	}
	password, err := util.EncryptPassword(user.Password)
	if err != nil {
		return "", nil
	}
	user.Password = password
	restErr := repository.SaveUser(&user)
	if restErr != nil {
		return "", restErr
	}
	return fmt.Sprintf("user %s created", user.Username), nil
}

func SignIn(user model.User) (string, *util.RestError) {
	var dbUser model.User
	if res := database.DBInstance.Where("username = ?", user.Username).First(&dbUser); res.RowsAffected == 0 {
		return "", util.NewRestErrBadRequest("Invalid username")
	}
	if !util.ComparePasswords(dbUser.Password, user.Password) {
		return "", util.NewRestErrBadRequest("Invalid password")
	}
	token, err := util.CreateToken(dbUser)
	if err != nil {
		return "", err
	}
	return token, nil
}
