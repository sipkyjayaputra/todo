package usecase

import (
	"net/http"
	"todo/model/dto"
	"todo/model/entity"
	"todo/repository"
)

type TaskUsecase interface {
	GetTasks() dto.Response
	GetTask(taskId uint) dto.Response
	CreateTask(taskDetail dto.Task) dto.Response
	UpdateTask(taskId uint, taskDetail dto.Task) dto.Response
	DeleteTask(taskId uint) dto.Response
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func CreateTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: repo,
	}
}

func (u *taskUsecase) GetTasks() dto.Response {
	response, err := u.taskRepository.GetTasks()
	if err != nil {
		return dto.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Error:      err.Error(),
			Data:       nil,
		}
	}

	tasks := []dto.Task{}
	for _, data := range response {
		tasks = append(tasks, dto.Task{
			Id:         data.Id,
			TaskName:   data.TaskName,
			TaskStatus: data.TaskStatus,
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
		})
	}

	return dto.Response{
		StatusCode: http.StatusOK,
		Status:     "Ok",
		Error:      nil,
		Data:       tasks,
	}
}

func (u *taskUsecase) GetTask(taskId uint) dto.Response {
	response, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return dto.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Error:      err.Error(),
			Data:       nil,
		}
	}

	return dto.Response{
		StatusCode: http.StatusOK,
		Status:     "Ok",
		Error:      nil,
		Data:       response,
	}
}

func (u *taskUsecase) CreateTask(taskDetail dto.Task) dto.Response {
	err := u.taskRepository.CreateTask(entity.Task{
		TaskName:   taskDetail.TaskName,
		TaskStatus: taskDetail.TaskStatus,
	})
	if err != nil {
		return dto.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Error:      err.Error(),
			Data:       nil,
		}
	}

	return dto.Response{
		StatusCode: http.StatusCreated,
		Status:     "Created",
		Error:      nil,
		Data:       nil,
	}
}

func (u *taskUsecase) UpdateTask(taskId uint, taskDetail dto.Task) dto.Response {
	err := u.taskRepository.UpdateTask(taskId, entity.Task{
		TaskName:   taskDetail.TaskName,
		TaskStatus: taskDetail.TaskStatus,
	})
	if err != nil {
		return dto.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Error:      err.Error(),
			Data:       nil,
		}
	}

	return dto.Response{
		StatusCode: http.StatusOK,
		Status:     "Ok",
		Error:      nil,
		Data:       nil,
	}
}

func (u *taskUsecase) DeleteTask(taskId uint) dto.Response {
	err := u.taskRepository.DeleteTask(taskId)
	if err != nil {
		return dto.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Error:      err.Error(),
			Data:       nil,
		}
	}

	return dto.Response{
		StatusCode: http.StatusOK,
		Status:     "Ok",
		Error:      nil,
		Data:       nil,
	}
}
