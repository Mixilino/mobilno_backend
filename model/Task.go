package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
}

type TaskDto struct {
	ID        uint   `json:"ID"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}

func (t *Task) ConvertToDto() *TaskDto {
	return &TaskDto{
		ID:        t.ID,
		Name:      t.Name,
		Completed: t.Completed,
		UserID:    t.UserID,
	}
}
func ConvertTaskArrayToDto(tasks []Task) []TaskDto {
	var tasksDto []TaskDto
	for _, task := range tasks {
		tasksDto = append(tasksDto, *task.ConvertToDto())
	}
	return tasksDto
}
