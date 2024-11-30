package services

import (
	"fmt"
	"gopoc/config/logger"
	"gopoc/models"
	"sync"
	"time"
)

var (
	taskDb []models.Task = make([]models.Task, 0)
	mu     sync.RWMutex
)

func GetTask(taskId int) (*models.Task, error) {
	logger.Trace(`Entering TaskService.GetTask`)
	mu.RLock()
	defer mu.RUnlock()
	for i := range taskDb {
		if taskDb[i].Id != 0 && taskDb[i].Id == taskId {
			logger.Info(`Task found, returning task as response`)
			return &taskDb[i], nil
		}
	}
	logger.Warn(`Task not found, returning error response`)
	logger.Trace(`Exitting TaskService.GetTask`)
	return nil, fmt.Errorf("task with id %d not found", taskId)
}

func CreateTask(task models.Task) (models.Task, error) {
	logger.Trace(`Entering TaskService.CreateTask`)
	mu.RLock()
	var newTaskId int = len(taskDb) + 1
	mu.RUnlock()
	task.Id = newTaskId
	task.CreatedOn = time.Now()
	logger.Info(`Appending upon available tasks`)
	mu.Lock()
	taskDb = append(taskDb, task)
	mu.Unlock()
	logger.Info(`Returning task as response`)
	logger.Trace(`Exitting TaskService.CreateTask`)
	return task, nil
}

func UpdateTask(task models.Task) (models.Task, error) {
	logger.Trace(`Entering TaskService.UpdateTask`)
	if task.Id != 0 {
		var dbTask *models.Task
		var err error
		logger.Info(`Fetching task using ID`)
		dbTask, err = GetTask(task.Id)
		if err != nil {
			logger.Error(`Error on finding task`, err)
			logger.Trace(`Exitting TaskService.UpdateTask`)
			return models.Task{}, err
		}
		logger.Info(`Updating Task details`)
		mu.RLock()
		dbTask.Id = task.Id
		dbTask.Title = task.Title
		dbTask.Description = task.Description
		dbTask.Status = task.Status
		mu.RUnlock()
		logger.Trace(`Exitting TaskService.UpdateTask`)
		return *dbTask, nil
	} else {
		logger.Warn(`Task to update not found`)
		logger.Trace(`Exitting TaskService.UpdateTask`)
		return models.Task{}, fmt.Errorf(`invalid identifier`)
	}
}

func DeleteTask(taskId int) (bool, error) {
	logger.Trace(`Entering TaskService.DeleteTask`)
	mu.RLock()
	defer mu.RUnlock()
	for index, task := range taskDb {
		if task.Id != 0 && task.Id == taskId {
			logger.Info(`Removing task`)
			taskDb = append(taskDb[:index], taskDb[index+1:]...)
			logger.Trace(`Exitting TaskService.DeleteTask`)
			return true, nil
		}
	}
	logger.Error(`Task ID to delete not found`)
	logger.Trace(`Exitting TaskService.DeleteTask`)
	return false, fmt.Errorf(`id not found on db`)
}

func GetAllTask() ([]models.Task, error) {
	logger.Trace(`Entering TaskService.GetAllTask`)
	logger.Info(`Fetching all Tasks`)
	mu.RLock()
	var resp = taskDb
	mu.RUnlock()
	logger.Trace(`Exitting TaskService.GetAllTask`)
	return resp, nil
}

func InvalidateTasks() {
	logger.Trace(`Entering TaskService.InvalidateTasks`)
	logger.Info(`Fetching all Tasks for status update`)
	mu.RLock()
	defer mu.RUnlock()
	for i := range taskDb {
		if taskDb[i].Status != "completed" && time.Now().Add(time.Hour*-1).After(taskDb[i].CreatedOn) {
			logger.Info(`Task found, changing task status as critical`)
			taskDb[i].Status = "critical"
		}
	}
	logger.Trace(`Exitting TaskService.InvalidateTasks`)
}
