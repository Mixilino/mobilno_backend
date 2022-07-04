package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
	"gorm.io/gorm"
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

func DeleteTask(taskID, userID uint) *util.RestError {
	res := database.DBInstance.Delete(&model.Task{
		Model:  gorm.Model{ID: taskID},
		UserID: userID,
	})
	if res.RowsAffected == 0 {
		return util.NewRestErrBadRequest("No task with that id")
	}
	return nil
}

func ModifyTask(task *model.Task, modifyName bool) *util.RestError {
	if !modifyName {
		oldTask := &model.Task{}
		oldTask.ID = task.ID
		database.DBInstance.Find(oldTask)
		task.Name = oldTask.Name
	}
	res := database.DBInstance.Save(task)
	if res.RowsAffected == 0 {
		return util.NewRestErrConflict("Task cant be created")
	}
	return nil
}

func GetTask(task *model.Task) *util.RestError {
	res := database.DBInstance.Find(task)
	if res.RowsAffected == 0 {
		return util.NewRestErrBadRequest("No task")
	}
	return nil
}
