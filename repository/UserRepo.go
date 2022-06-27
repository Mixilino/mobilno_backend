package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
)

func SaveUser(user *model.User) *util.RestError {
	res := database.DBInstance.Create(user)
	if res.RowsAffected == 0 {
		return util.NewRestErrConflict("username already exists")
	}
	return nil
}
