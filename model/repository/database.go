package repository

import (
	"errors"
	"go-task-api/model/entity"
	"math"
	"sync"
)

type SerialNumber struct {
	id int
}

func (sn *SerialNumber) Init() {
	sn.id = 0
}

func (sn *SerialNumber) NextID() (id int) {
	id = sn.id
	if sn.id == math.MaxInt {
		sn.id = 0
	} else {
		sn.id++
	}
	return
}

type MockDB struct {
	tasksMap map[int]entity.TaskEntity
	node     SerialNumber
	lock     sync.Mutex
}

func (db *MockDB) Init() {
	db.node.Init()
	db.tasksMap = make(map[int]entity.TaskEntity)
}

func (db *MockDB) GetTasks() (tasks []entity.TaskEntity, err error) {
	for _, val := range db.tasksMap {
		tasks = append(tasks, val)
	}
	return
}

func (db *MockDB) InsertTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	db.lock.Lock()
	defer db.lock.Unlock()
	task.ID = db.node.NextID()
	db.tasksMap[task.ID] = task
	result = db.tasksMap[task.ID]
	return
}

func (db *MockDB) UpdateTask(task entity.TaskEntity) (result entity.TaskEntity, err error) {
	db.lock.Lock()
	defer db.lock.Unlock()
	if val, ok := db.tasksMap[task.ID]; !ok {
		err = errors.New("task id is not exist")
	} else {
		val.Status = task.Status
		val.Name = task.Name
		db.tasksMap[task.ID] = val
		result = val
	}
	return
}

func (db *MockDB) DeleteTask(id int) (err error) {
	db.lock.Lock()
	defer db.lock.Unlock()
	if _, ok := db.tasksMap[id]; !ok {
		err = errors.New("task id is not exist")
	} else {
		delete(db.tasksMap, id)
	}

	return
}

var Db *MockDB

func init() {
	Db = &MockDB{}
	Db.Init()
}
