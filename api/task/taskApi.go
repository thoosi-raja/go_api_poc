package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gopoc/config/logger"
	HttpConstants "gopoc/constants/http_constants"

	"gopoc/models"
	"gopoc/services"

	"github.com/go-chi/chi/v5"
)

func init() {
	logger.Info(`Initializing Task API`)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	logger.Trace("Entering /task/GetTask")
	taskIdStr := chi.URLParam(r, "taskId")
	if taskIdStr == "" {
		services.ApiErrorResponse(w,
			`taskId query parameter is required`,
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/GetTask")
		return
	}
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		services.ApiErrorResponse(w,
			`taskId query parameter format is incorrect`,
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/GetTask")
		return
	}
	resultChannel, errChannel := services.AsyncFunctionWith1Param(services.GetTask, taskId)
	select {
	case task := <-resultChannel:
		{
			taskJSON, err := json.Marshal(task)
			if err != nil {
				services.ApiErrorResponse(w,
					"Failed to marshal task to JSON",
					&HttpConstants.INTERNAL_SERVER_ERROR)
				logger.Trace("Exitting /task/GetTask")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(taskJSON)
			logger.Trace("Exitting /task/GetTask")
			return
		}
	case err := <-errChannel:
		{
			services.ApiErrorResponse(w,
				err.Error(),
				&HttpConstants.INTERNAL_SERVER_ERROR)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Trace("Exitting /task/GetTask")
			return
		}
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	logger.Trace("Entering /task/createTask")
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		services.ApiErrorResponse(w,
			fmt.Sprintf("Error decoding task: %v", err),
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/createTask")
		return
	}
	resultChannel, errChannel := services.AsyncFunctionWith1Param(services.CreateTask, task)
	select {
	case updatedTask := <-resultChannel:
		{
			taskJSON, err := json.Marshal(updatedTask)
			if err != nil {
				services.ApiErrorResponse(w,
					"Failed to marshal task to JSON",
					&HttpConstants.INTERNAL_SERVER_ERROR)
				logger.Trace("Exitting /task/createTask")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(taskJSON)
			logger.Trace("Exitting /task/createTask")
		}
	case err2 := <-errChannel:
		{
			services.ApiErrorResponse(w,
				fmt.Sprintf("Error: %v", err2),
				&HttpConstants.INTERNAL_SERVER_ERROR)
			logger.Trace("Exitting /task/createTask")
			return
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	logger.Trace("Entering /task/UpdateTask")
	taskIdStr := chi.URLParam(r, "taskId")
	if taskIdStr == "" {
		services.ApiErrorResponse(w,
			"taskId query parameter is required",
			&HttpConstants.BAD_REQUEST)
		http.Error(w, "taskId query parameter is required", http.StatusBadRequest)
		logger.Trace("Exitting /task/GetTask")
		return
	}
	taskId, err2 := strconv.Atoi(taskIdStr)
	if err2 != nil {
		services.ApiErrorResponse(w,
			fmt.Sprintf("Error in query paramter format taskId: %v", err2),
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/UpdateTask")
		return
	}
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		services.ApiErrorResponse(w,
			fmt.Sprintf("Error decoding task: %v", err),
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/UpdateTask")
		return
	}
	if taskId != task.Id {
		services.ApiErrorResponse(w,
			fmt.Sprintf("TaskId not matching: %v", err),
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/UpdateTask")
		return
	}
	resultChannel, errChannel := services.AsyncFunctionWith1Param(services.UpdateTask, task)
	select {
	case updatedTask := <-resultChannel:
		{
			taskJSON, err := json.Marshal(updatedTask)
			if err != nil {
				services.ApiErrorResponse(w,
					"Failed to marshal task to JSON",
					&HttpConstants.INTERNAL_SERVER_ERROR)
				logger.Trace("Exitting /task/UpdateTask")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(taskJSON)
			logger.Trace("Exitting /task/UpdateTask")
		}
	case err2 := <-errChannel:
		{
			services.ApiErrorResponse(w,
				fmt.Sprintf("Error: %v", err2),
				&HttpConstants.INTERNAL_SERVER_ERROR)
			logger.Trace("Exitting /task/UpdateTask")
			return
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	logger.Trace("Entering /task/DeleteTask")
	taskIdStr := chi.URLParam(r, "taskId")
	if taskIdStr == "" {
		services.ApiErrorResponse(w,
			"taskId query parameter is required",
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/DeleteTask")
		return
	}
	taskId, err2 := strconv.Atoi(taskIdStr)
	if err2 != nil {
		services.ApiErrorResponse(w,
			fmt.Sprintf("Error in query paramter format taskId: %v", err2),
			&HttpConstants.BAD_REQUEST)
		logger.Trace("Exitting /task/DeleteTask")
		return
	}
	resultChannel, errChannel := services.AsyncFunctionWith1Param(services.DeleteTask, taskId)
	select {
	case resp := <-resultChannel:
		{
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%t", resp)))
			logger.Trace("Exitting /task/DeleteTask")
		}
	case err2 := <-errChannel:
		{
			services.ApiErrorResponse(w,
				err2.Error(),
				&HttpConstants.INTERNAL_SERVER_ERROR)
			logger.Trace("Exitting /task/DeleteTask")
			return
		}
	}
}

func GetAlltask(w http.ResponseWriter, r *http.Request) {
	logger.Trace("Entering /task/GetAlltask")
	resultChannel, errChannel := services.AsyncFunctionWithoutParam(services.GetAllTask)
	select {
	case tasks := <-resultChannel:
		{
			taskJSON, err := json.Marshal(tasks)
			if err != nil {
				services.ApiErrorResponse(w,
					"Failed to marshal task to JSON",
					&HttpConstants.INTERNAL_SERVER_ERROR)
				logger.Trace("Exitting /task/UpdateTask")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(taskJSON)
			logger.Trace("Exitting /task/GetAlltask")
		}
	case err2 := <-errChannel:
		{
			services.ApiErrorResponse(w,
				err2.Error(),
				&HttpConstants.INTERNAL_SERVER_ERROR)
			logger.Trace("Exitting /task/GetAlltask")
			return
		}
	}
}
