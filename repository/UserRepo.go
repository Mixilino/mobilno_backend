package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
	"fmt"
)

func SaveUser(user *model.User) *util.RestError {
	res := database.DBInstance.Create(user)
	fmt.Println(res)
	if res.RowsAffected == 0 {
		return util.NewRestErrConflict("username already exists")
	}
	return nil
}
