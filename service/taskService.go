package service

import (
	"backend/model"
	"backend/repository"
	"backend/util"
	"gorm.io/gorm"
)

func CreateTask(task *model.Task) *util.RestError {
	if task.Name == "" {
		return util.NewRestErrBadRequest("Task name cannot be empty")
	}
	task.Completed = false
	if err := repository.CreateTask(task); err != nil {
		return err
	}
	return nil
}

func GetAllTasks(userID uint) ([]model.Task, *util.RestError) {
	tasks, err := repository.GetAllTasks(userID)
	if err != nil {
		return nil, util.NewRestErrInternalServerError()
	}
	return tasks, nil
}

func DeleteTask(taskID, userID uint) *util.RestError {
	return repository.DeleteTask(taskID, userID)
}

func ModifyTask(task *model.Task) *util.RestError {
	modifyName := task.Name == ""
	return repository.ModifyTask(task, !modifyName)
}
func GetTask(userID, taskId uint) (*model.TaskDto, *util.RestError) {
	task := &model.Task{
		Model:  gorm.Model{ID: taskId},
		UserID: userID,
	}
	err := repository.GetTask(task)
	if err != nil {
		return nil, err
	}
	return task.ConvertToDto(), nil
}
