package services

import (
	"encoding/json"
	"gopoc/models"
	"net/http"
	"time"
)

func ApiErrorResponse(
	writer http.ResponseWriter,
	message string,
	statusCode *int) {
	if statusCode == nil {
		defaultStatus := http.StatusInternalServerError
		statusCode = &defaultStatus
	}
	apiError := models.ApiError{
		Timestamp:  time.Now(),
		Message:    message,
		StatusCode: *statusCode,
	}
	writer.WriteHeader(*statusCode)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(apiError)
}
