package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"onehub/crons"
	"onehub/fs"
	"onehub/models"
	"time"
)

func ApiPutRPCSecret(c *gin.Context) {
	var request struct{
		Secret string `json:"password"`
	}
	_ = c.BindJSON(&request)
	viper.Set("RPCSecret", request.Secret)
	viper.WriteConfig()
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "Secret set.",
		Data:    nil,
	})
}

func ApiGetStats(c *gin.Context) {
	entry := crons.Cron.Entries()[0]
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    map[string]interface{}{
			"files_count": fs.Root.FilesCount(),
			"folders_count": fs.Root.FoldersCount(),
			"next_schedule_update": entry.Next,
		},
	})
}

func ApiCheckPassword(c *gin.Context) {
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    nil,
	})
}

type ConfigRes struct {
	RefreshInterval int `json:"refresh_interval"`
	NextScheduleUpdate time.Time `json:"next_schedule_update"`
	Title string `json:"title"`
}

func ApiAdminGetConfig(c *gin.Context) {
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    ConfigRes{
			RefreshInterval: viper.GetInt("RefreshInterval"),
			NextScheduleUpdate: crons.Cron.Entry(crons.FsCronID).Next,
			Title: viper.GetString("Title"),
		},
	})
}

func ApiAdminPutConfig(c *gin.Context) {
	var request ConfigRes
	_ = c.BindJSON(&request)

	if request.RefreshInterval < 5 || request.RefreshInterval > 60 {
		c.JSON(200, models.OkJson{
			Success: false,
			Message: "刷新间隔应不少于5分钟且不大于60分钟",
			Data:    ConfigRes{
				RefreshInterval: viper.GetInt("RefreshInterval"),
				NextScheduleUpdate: crons.Cron.Entry(crons.FsCronID).Next,
				Title: viper.GetString("Title"),
			},
		})
		return
	}

	viper.Set("RefreshInterval", request.RefreshInterval)
	viper.Set("Title", request.Title)
	viper.WriteConfig()
	crons.RestartCron()
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    nil,
	})
}
