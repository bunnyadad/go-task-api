package repository

import (
	"go-task-api/model/entity"
)

type TaskRepository interface {
	GetTasks() (tasks []entity.TaskEntity, err error)
	InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error)
	UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error)
	DeleteTask(id int) (err error)
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (tr *taskRepository) GetTasks() (tasks []entity.TaskEntity, err error) {
	tasks, err = Db.GetTasks()
	return
}

func (tr *taskRepository) InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	result, err = Db.InsertTask(task)
	return
}

func (tr *taskRepository) UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	result, err = Db.UpdateTask(task)
	return
}

func (tr *taskRepository) DeleteTask(id int) (err error) {
	err = Db.DeleteTask(id)
	return
}
