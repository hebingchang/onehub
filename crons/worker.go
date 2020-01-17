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
var FsCronID cron.EntryID
var TokenCronID cron.EntryID

func StartCron() {
	Cron = cron.New()
	FsCronID, _ = Cron.AddFunc("*/"+viper.GetString("RefreshInterval")+" * * * *", func() {
		fs.Root = workers.WalkDrive()
		fs.LastUpdate = time.Now()
		fs.WriteFsDataFile()
	})
	TokenCronID, _ = Cron.AddFunc("*/30 * * * *", func() {
		services.RefreshToken()
	})
	Cron.Start()
}

func RestartCron() {
	Cron.Stop()
	StartCron()
}