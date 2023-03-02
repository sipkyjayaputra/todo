package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id         uint           `json:"id"`
	TaskName   string         `json:"task_name"`
	TaskStatus int            `json:"task_status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (Task) TableName() string {
	return "task_list"
}
