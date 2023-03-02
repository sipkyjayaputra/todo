package repository

import (
	"time"
	"todo/model/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks() ([]entity.Task, error)
	GetTask(taskId uint) (entity.Task, error)
	CreateTask(taskDetail entity.Task) error
	UpdateTask(taskId uint, taskDetail entity.Task) error
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	DB *gorm.DB
}

func CreateTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (r *taskRepository) GetTasks() ([]entity.Task, error) {
	tasks := []entity.Task{}
	err := r.DB.Model(&entity.Task{}).Scan(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (r *taskRepository) GetTask(taskId uint) (entity.Task, error) {
	task := entity.Task{}
	if err := r.DB.Model(&entity.Task{}).Where("id", taskId).Scan(&task).Error; err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) CreateTask(taskDetail entity.Task) error {
	taskDetail.CreatedAt = time.Now()
	taskDetail.UpdatedAt = time.Now()
	if err := r.DB.Model(&entity.Task{}).Create(&taskDetail).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) UpdateTask(taskId uint, taskDetail entity.Task) error {
	taskDetail.UpdatedAt = time.Now()
	if err := r.DB.Model(&entity.Task{}).Where("id = ?", taskId).Updates(map[string]interface{}{"task_name": taskDetail.TaskName, "task_status": taskDetail.TaskStatus, "updated_at": taskDetail.UpdatedAt}).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) DeleteTask(taskId uint) error {
	if err := r.DB.Model(&entity.Task{}).Delete(&entity.Task{Id: taskId}).Error; err != nil {
		return err
	}
	return nil
}
