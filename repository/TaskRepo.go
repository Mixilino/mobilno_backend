package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
)

func CreateTask(task *model.Task) *util.RestError {
	res := database.DBInstance.Create(task)
	if res.RowsAffected == 0 {
		return util.NewRestErrConflict("Task cant be created")
	}
	return nil
}

func GetAllTasks(userID uint) ([]model.Task, *util.RestError) {
	var tasks []model.Task
	database.DBInstance.Where("user_id = ?", userID).Find(&tasks)
	return tasks, nil
}
