package test

import (
	"errors"
	"net/http"

	"go-task-api/model/entity"

	"github.com/gin-gonic/gin"
)

type MockTaskController struct {
}

func (mtc *MockTaskController) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (mtc *MockTaskController) PostTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}

func (mtc *MockTaskController) PutTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (mtc *MockTaskController) DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

type MockTaskRepository struct {
}

func (mtr *MockTaskRepository) GetTasks() (tasks []entity.TaskEntity, err error) {
	tasks = []entity.TaskEntity{}
	return
}

func (mtr *MockTaskRepository) InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	result = task
	return
}

func (mtr *MockTaskRepository) UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	return
}

func (mtr *MockTaskRepository) DeleteTask(id int) (err error) {
	return
}

type MockTaskRepositoryGetTasksExist struct {
}

func (mtrgex *MockTaskRepositoryGetTasksExist) GetTasks() (tasks []entity.TaskEntity, err error) {
	tasks = []entity.TaskEntity{}
	tasks = append(tasks, entity.TaskEntity{ID: 1, Name: "title1", Status: 0})
	tasks = append(tasks, entity.TaskEntity{ID: 2, Name: "title2", Status: 0})
	return
}

func (mtrgex *MockTaskRepositoryGetTasksExist) InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	return
}

func (mtrgex *MockTaskRepositoryGetTasksExist) UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	return
}

func (mtrgex *MockTaskRepositoryGetTasksExist) DeleteTask(id int) (err error) {
	return
}

type MockTaskRepositoryError struct {
}

func (mtrgtn *MockTaskRepositoryError) GetTasks() (tasks []entity.TaskEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgie *MockTaskRepositoryError) InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgue *MockTaskRepositoryError) UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgde *MockTaskRepositoryError) DeleteTask(id int) (err error) {
	err = errors.New("unexpected error occurred")
	return
}
