package services

import (
	"gopoc/config/logger"
)

func TaskStatusCheckCron() {
	logger.Trace("Entering TaskStatusCheckCron")
	InvalidateTasks()
}
