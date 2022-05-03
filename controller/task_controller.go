package controller

import (
	"net/http"
	"strconv"

	"go-task-api/controller/rest"
	"go-task-api/model/entity"
	"go-task-api/model/repository"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetTasks(c *gin.Context)
	PostTask(c *gin.Context)
	PutTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type taskController struct {
	tr repository.TaskRepository
}

func NewTaskController(tr repository.TaskRepository) TaskController {
	return &taskController{tr}
}

func (tc *taskController) GetTasks(c *gin.Context) {
	tasks, err := tc.tr.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var taskResponses []rest.TaskResponse
	for _, v := range tasks {
		taskResponses = append(taskResponses, rest.TaskResponse{Id: v.ID, Name: v.Name, Status: v.Status})
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"result": taskResponses})
}

func (tc *taskController) PostTask(c *gin.Context) {
	var taskRequest rest.TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := entity.TaskEntity{Name: taskRequest.Name}
	result, err := tc.tr.InsertTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tasksResponse rest.TasksResponse
	tasksResponse.Tasks = append(tasksResponse.Tasks, rest.TaskResponse{Id: result.ID, Name: result.Name, Status: result.Status})
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"result": tasksResponse.Tasks})
}

func (tc *taskController) PutTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not correct!"})
		return
	}
	var taskRequest rest.TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task := entity.TaskEntity{ID: taskId, Name: taskRequest.Name, Status: taskRequest.Status}
	result, err := tc.tr.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tasksResponse rest.TasksResponse
	tasksResponse.Tasks = append(tasksResponse.Tasks, rest.TaskResponse{Id: result.ID, Name: result.Name, Status: result.Status})
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"result": tasksResponse.Tasks})
}

func (tc *taskController) DeleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not correct!"})
		return
	}

	err = tc.tr.DeleteTask(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
