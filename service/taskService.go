package service

import (
	"backend/model"
	"backend/repository"
	"backend/util"
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

func GetAllTasks(userId uint) ([]model.Task, *util.RestError) {
	tasks, err := repository.GetAllTasks(userId)
	if err != nil {
		return nil, util.NewRestErrInternalServerError()
	}
	return tasks, nil
}
