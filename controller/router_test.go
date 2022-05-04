package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"go-task-api/test"

	"github.com/gin-gonic/gin"
)

var mux *gin.Engine

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := NewRouter(&test.MockTaskController{})
	mux = target.HandleTasksRequest()
}

func TestGetTasks(t *testing.T) {
	r, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostTask(t *testing.T) {
	json := strings.NewReader(`{"name":"test"}`)
	r, _ := http.NewRequest("POST", "/task", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTask(t *testing.T) {
	json := strings.NewReader(`{"name":"test","status":1}`)
	r, _ := http.NewRequest("PUT", "/task/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/task/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/task", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Response cod is %v", w.Code)
	}
}
