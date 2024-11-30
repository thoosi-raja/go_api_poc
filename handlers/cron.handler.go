package handlers

import (
	"gopoc/services"

	"github.com/robfig/cron/v3"
)

func CronHandler() {
	cron := cron.New()
	cron.AddFunc("@every 1m", services.TaskStatusCheckCron)
	cron.Start()
}
