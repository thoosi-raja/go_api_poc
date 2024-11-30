package test

import (
	"gopoc/models"
	"gopoc/services"
	"testing"
	"time"
)

func TestGetAllTask(t *testing.T) {
	result, err := services.GetAllTask()
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Error("Error Task")
	}
}

func TestCreateTask(t *testing.T) {
	newTask := models.Task{
		Id:          1,
		Title:       "title 1",
		Description: "description 1",
		Status:      "",
		CreatedOn:   time.Now(),
	}
	result, err := services.CreateTask(newTask)
	if err != nil {
		t.Error(err)
	}
	if result.Id == 0 {
		t.Error("Task created with invalid ID ")
	}
	if result.Title != newTask.Title {
		t.Error("Task created with invalid Title ")
	}
	if result.Description != newTask.Description {
		t.Error("Task created with invalid Description ")
	}
	if result.Status != newTask.Status {
		t.Error("Task created with invalid Status ")
	}
}

func TestGetAllTaskAfterCreation(t *testing.T) {
	newTask := models.Task{
		Id:          1,
		Title:       "title 1",
		Description: "description 1",
		Status:      "",
		CreatedOn:   time.Now(),
	}
	services.CreateTask(newTask)
	result, err := services.GetAllTask()
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("Invalid Tasks found")
		return
	}
	if result[0].Id == 0 {
		t.Error("Task created with invalid ID ")
	}
	if result[0].Title != newTask.Title {
		t.Error("Task created with invalid Title ")
	}
	if result[0].Description != newTask.Description {
		t.Error("Task created with invalid Description ")
	}
	if result[0].Status != newTask.Status {
		t.Error("Task created with invalid Status ")
	}
}

func TestUpdateTask(t *testing.T) {
	newTask := models.Task{
		Id:          1,
		Title:       "title 1",
		Description: "description 1",
		Status:      "",
		CreatedOn:   time.Now(),
	}
	services.CreateTask(newTask)
	updateTask := models.Task{
		Id:          1,
		Title:       "updated title 1",
		Description: "updated description 1",
		Status:      "updated status",
		CreatedOn:   time.Now(),
	}
	result, err := services.UpdateTask(updateTask)
	if err != nil {
		t.Error(err)
	}
	if result.Id == 0 {
		t.Error("Task created with invalid ID ")
	}
	if result.Title != updateTask.Title {
		t.Error("Task created with invalid Title ")
	}
	if result.Description != updateTask.Description {
		t.Error("Task created with invalid Description ")
	}
	if result.Status != updateTask.Status {
		t.Error("Task created with invalid Status ")
	}
}

func TestDeleteTask(t *testing.T) {
	newTask := models.Task{
		Id:          1,
		Title:       "title 1",
		Description: "description 1",
		Status:      "",
		CreatedOn:   time.Now(),
	}
	services.CreateTask(newTask)
	services.DeleteTask(newTask.Id)
	result, err := services.GetAllTask()
	if err != nil {
		t.Error(err)
	}
	if len(result) != 0 {
		t.Error("Task failed to delete")
	}
}

func TestCronStatusUpdateTask(t *testing.T) {
	newTask := models.Task{
		Id:          1,
		Title:       "title 1",
		Description: "description 1",
		Status:      "in-p",
		CreatedOn:   time.Now(),
	}
	services.CreateTask(newTask)
	time.Sleep(1 * time.Hour)
	services.InvalidateTasks()
	result, err := services.GetAllTask()
	if err != nil {
		t.Error(err)
	}
	if result[0].Status != "critical" {
		t.Error("Task failed to update status to critical")
	}
}
