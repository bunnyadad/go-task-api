package controller

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	HandleTasksRequest() *gin.Engine
}

type router struct {
	tc TaskController
}

func NewRouter(tc TaskController) Router {
	return &router{tc}
}

func (ro *router) HandleTasksRequest() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
	})
	r.GET("/tasks", ro.tc.GetTasks)
	r.POST("/task", ro.tc.PostTask)
	r.PUT("/task/:id", ro.tc.PutTask)
	r.DELETE("task/:id", ro.tc.DeleteTask)
	return r
}
