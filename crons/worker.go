package crons

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"onehub/fs"
	"onehub/services"
	"onehub/workers"
	"time"
)

var Cron *cron.Cron

func StartCron() {
	Cron = cron.New()
	Cron.AddFunc("*/" + viper.GetString("RefreshInterval") + " * * * *", func() {
		services.RefreshToken()
		fs.Root = workers.WalkDrive()
		fs.LastUpdate = time.Now()
		fs.WriteFsDataFile()
	})
	Cron.Start()
}

func RestartCron() {
	Cron.Stop()
	StartCron()
}