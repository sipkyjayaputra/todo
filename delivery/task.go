package delivery

import (
	"net/http"
	"strconv"
	"todo/model/dto"
	"todo/usecase"

	"github.com/gin-gonic/gin"
)

type TaskDelivery interface {
	GetTasks(*gin.Context)
	GetTask(*gin.Context)
	CreateTask(*gin.Context)
	UpdateTask(*gin.Context)
	DeleteTask(*gin.Context)
}

type taskDelivery struct {
	taskUsecase usecase.TaskUsecase
}

func CreateTaskDelivery(uc usecase.TaskUsecase) TaskDelivery {
	return &taskDelivery{
		taskUsecase: uc,
	}
}

func (t *taskDelivery) GetTasks(c *gin.Context) {
	response := t.taskUsecase.GetTasks()

	c.JSON(response.StatusCode, response)
}

func (t *taskDelivery) GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	response := t.taskUsecase.GetTask(uint(id))
	c.JSON(response.StatusCode, response)
}

func (t *taskDelivery) CreateTask(c *gin.Context) {
	request := dto.Task{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Error bad request")
	}

	response := t.taskUsecase.CreateTask(request)
	c.JSON(response.StatusCode, response)
}

func (t *taskDelivery) UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	request := dto.Task{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Error bad request")
	}

	response := t.taskUsecase.UpdateTask(uint(id), request)
	c.JSON(response.StatusCode, response)
}

func (t *taskDelivery) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	response := t.taskUsecase.DeleteTask(uint(id))
	c.JSON(response.StatusCode, response)
}
