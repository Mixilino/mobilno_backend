package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}
