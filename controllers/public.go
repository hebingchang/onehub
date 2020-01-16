package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"onehub/fs"
	"onehub/models"
)

func ApiGetConfig(c *gin.Context) {
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    map[string]interface{}{
			"title": viper.GetString("Title"),
			"last_update": fs.LastUpdate,
		},
	})
}

func HasConfig() bool {
	return viper.GetString("RPCSecret") != "" &&
		viper.GetString("DriveID") != "" &&
		viper.GetString("ItemID") != ""
}

func ApiHasConfig(c *gin.Context) {
	if HasConfig() {
		c.JSON(200, models.OkJson{
			Success: true,
			Message: "ok",
			Data:    true,
		})
	} else {
		c.JSON(200, models.OkJson{
			Success: true,
			Message: "ok",
			Data:    false,
		})
	}
}

func ApiGetTree(c *gin.Context) {
	tree := fs.Root.Tree()
	tree.ID = "/"
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    tree,
	})
}

func ApiGetChild(c *gin.Context) {
	var request struct{
		Path string `json:"path"`
	}
	_ = c.BindJSON(&request)
	data, err := fs.GetChildren(request.Path)
	if err != nil {
		c.JSON(400, models.ErrorJson{
			Success: false,
			Message: err.Error(),
		})
	} else {
		c.JSON(200, models.OkJson{
			Success: true,
			Message: "ok",
			Data:    data,
		})
	}
}

func ApiDownloadItem(c *gin.Context) {
	item_id := c.Param("item_id")
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	req_url := "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/items/" + item_id
	r, err := req.Get(req_url, header)
	if err != nil {
		log.Println(err)
	}

	var item models.DriveItem
	_ = r.ToJSON(&item)
	c.Redirect(http.StatusMovedPermanently, item.MicrosoftGraphDownloadURL)
}

func ApiGetItem(c *gin.Context) {
	item_id := c.Param("item_id")
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	req_url := "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/items/" + item_id
	r, err := req.Get(req_url, header)
	if err != nil {
		log.Println(err)
	}

	var item interface{}
	_ = r.ToJSON(&item)
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    item,
	})
}