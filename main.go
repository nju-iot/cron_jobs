package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/nju-iot/cron_jobs/caller"
	"github.com/nju-iot/cron_jobs/config"
	"github.com/nju-iot/cron_jobs/cron_action"
	"github.com/nju-iot/cron_jobs/logs"
)

func main() {
	config.LoadConfig()
	logs.InitLogs()
	caller.InitClient()

	InitCronLoader()
}

func InitCronLoader() {

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.TagsUnique()

	// your code ...
	scheduler.Every(10).Seconds().Tag("update_edgex_status").Do(cron_action.UpdateEdgexStatus)

	scheduler.StartAsync()
	scheduler.StartBlocking()
}
