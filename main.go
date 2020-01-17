package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"onehub/controllers"
	"onehub/crons"
	"onehub/middlewares"
	"onehub/models"
	"os"
	"time"
)

func readConfig() {
	viper.SetDefault("Token", models.Token{})
	viper.SetDefault("RPCSecret", "")
	viper.SetDefault("DriveID", "")
	viper.SetDefault("Title", "OneHub")
	viper.SetDefault("ItemID", "")
	viper.SetDefault("RefreshInterval", 30)
	viper.SetDefault("ClientID", "5ff35b5b-320b-441d-846e-baf6f2dce255")
	viper.SetDefault("ClientSecret", "BoplO[9-yeXZdPuiC:3x9soZ?RC9ilMs")

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/onehub/")
	viper.AddConfigPath("$HOME/.onehub")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, _ = os.Create("./config/config.yaml")
		}
	}
}

func init() {
	readConfig()
}

func main() {
	r := gin.Default()
	r.Use(location.New(location.Config{
		Host:   "localhost:8080",
		Scheme: "http",
		Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-Host"},
	}))
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "token"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.LoadHTMLGlob("templates/*")
	r.Static("/drive", "./frontend/public/dist")
	r.Static("/admin", "./frontend/admin/dist")

	r.GET("/", func(c *gin.Context) {
		if controllers.HasConfig() {
			c.Redirect(http.StatusMovedPermanently, "/drive")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/admin/#/init")
		}
	})

	/* controllers.RefreshToken()
	fs.Root = crons.WalkDrive()
	fs.LastUpdate = time.Now()
	fs.WriteFsDataFile() */

	crons.StartCron()

	v1 := r.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.Use(middlewares.AdminMiddleware)
			admin.GET("/oauth", controllers.OauthRedirect)
			admin.GET("/oauth/callback", controllers.OauthCallback)
			admin.PUT("/secret", controllers.ApiPutRPCSecret)
			admin.GET("/drives", controllers.ApiGetDrives)
			admin.PUT("/drive", controllers.ApiPutDrive)
			admin.POST("/drive-item", controllers.ApiGetDriveItem)
			admin.PUT("/drive-item", controllers.ApiPutDriveItem)
			admin.GET("/stats", controllers.ApiGetStats)
			admin.GET("/check-password", controllers.ApiCheckPassword)
			admin.GET("/profile", controllers.OauthProfile)
			admin.GET("/profile/drive", controllers.GetCurrentDrive)
			admin.GET("/profile/item", controllers.GetCurrentItem)
			admin.GET("/config", controllers.ApiAdminGetConfig)
			admin.PUT("/config", controllers.ApiAdminPutConfig)
		}

		public := v1.Group("/public")
		{
			public.GET("/has-config", controllers.ApiHasConfig)
			public.GET("/config", controllers.ApiGetConfig)
			public.POST("/files", controllers.ApiGetChild)
			public.GET("/tree", controllers.ApiGetTree)
			public.GET("/download/:item_id", controllers.ApiDownloadItem)
			public.GET("/item/:item_id", controllers.ApiGetItem)
		}
	}
	r.Run(":8080")
}
