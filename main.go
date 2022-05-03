package main

import (
	"go-task-api/controller"
	"go-task-api/model/repository"
)

var tr = repository.NewTaskRepository()
var tc = controller.NewTaskController(tr)
var ro = controller.NewRouter(tc)

func main() {
	r := ro.HandleTasksRequest()
	r.Run()
}
