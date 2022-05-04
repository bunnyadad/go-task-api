package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-task-api/controller/rest"
	"go-task-api/test"

	"github.com/gin-gonic/gin"
)

func TestGetTasks_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks", nil)

	target := NewTaskController(&test.MockTaskRepository{})
	target.GetTasks(c)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var tasksResponse rest.TasksResponse
	json.Unmarshal(body, &tasksResponse)
	if len(tasksResponse.Tasks) != 0 {
		t.Errorf("Response is %v", tasksResponse.Tasks)
	}
}

func TestGetTasks_ExistTask(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks", nil)

	target := NewTaskController(&test.MockTaskRepositoryGetTasksExist{})
	target.GetTasks(c)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var tasksResponse rest.TasksResponse
	json.Unmarshal(body, &tasksResponse)
	if len(tasksResponse.Tasks) != 2 {
		t.Errorf("Response is %v", tasksResponse.Tasks)
	}
}

func TestGetTasks_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks", nil)

	target := NewTaskController(&test.MockTaskRepositoryError{})
	target.GetTasks(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostTask_OK(t *testing.T) {
	jsonStr := strings.NewReader(`{"name":"買晚餐"}`)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/task/", jsonStr)

	target := NewTaskController(&test.MockTaskRepository{})
	target.PostTask(c)

	if w.Code != http.StatusCreated {
		t.Errorf("Response cod is %v", w.Code)
	}
	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var tasksResponse rest.TasksResponse
	json.Unmarshal(body, &tasksResponse)
	if len(tasksResponse.Tasks) == 0 {
		t.Errorf("Response is %v", tasksResponse.Tasks)
	}
}

func TestPostTask_Error(t *testing.T) {
	jsonStr := strings.NewReader(`{"name":"買晚餐"}`)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/task/", jsonStr)

	target := NewTaskController(&test.MockTaskRepositoryError{})
	target.PostTask(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != "" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPutTask_OK(t *testing.T) {
	jsonStr := strings.NewReader(`{"name":"買晚餐"}`)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/task/", jsonStr)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}
	target := NewTaskController(&test.MockTaskRepository{})
	target.PutTask(c)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTask_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/task/", nil)

	target := NewTaskController(&test.MockTaskRepository{})
	target.PutTask(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTask_Error(t *testing.T) {
	w := httptest.NewRecorder()
	jsonStr := strings.NewReader(`{"name":"買晚餐"}`)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/task/", jsonStr)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}

	target := NewTaskController(&test.MockTaskRepositoryError{})
	target.PutTask(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTask_OK(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/task/", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}

	target := NewTaskController(&test.MockTaskRepository{})
	target.DeleteTask(c)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTask_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/task/", nil)

	target := NewTaskController(&test.MockTaskRepositoryError{})
	target.DeleteTask(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTask_Error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/task/", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "2"}}

	target := NewTaskController(&test.MockTaskRepositoryError{})
	target.DeleteTask(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Response cod is %v", w.Code)
	}
}
