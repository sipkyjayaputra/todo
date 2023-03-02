package app

import (
	"todo/delivery"
	"todo/repository"
	"todo/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	taskRepository := repository.CreateTaskRepository(db)
	taskUsecase := usecase.CreateTaskUsecase(taskRepository)
	taskDelivery := delivery.CreateTaskDelivery(taskUsecase)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	{
		router.GET("/task", taskDelivery.GetTasks)
		router.GET("/task/:id", taskDelivery.GetTask)
		router.POST("/task", taskDelivery.CreateTask)
		router.PUT("/task/:id", taskDelivery.UpdateTask)
		router.DELETE("/task/:id", taskDelivery.DeleteTask)
	}

	return router
}
