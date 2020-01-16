package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"onehub/models"
	"strings"
)

func token() (models.Token) {
	var token models.Token
	_ = mapstructure.Decode(viper.Get("Token"), &token)
	return token
}

func ApiGetDrives(c *gin.Context) {
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	r, err := req.Get("https://graph.microsoft.com/v1.0/me/drives", header)
	if err != nil {
		log.Println(err)
	}
	var response models.DriveResponse
	_ = json.Unmarshal(r.Bytes(), &response)
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    response.Value,
	})
}

func ApiPutDrive(c *gin.Context) {
	var request struct{
		DriveID string `json:"drive_id"`
	}
	_ = c.BindJSON(&request)
	viper.Set("DriveID", request.DriveID)
	viper.Set("ItemID", "")
	viper.WriteConfig()
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "Drive ID set.",
		Data:    nil,
	})
}

func GetDriveItem(path string) ([]interface{}) {
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}

	var req_url string
	if path[len(path) - 1] == '/' {
		path = path[0:len(path)-1]
	}
	if path == "" {
		req_url = "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/root/children"
	} else if strings.Count(path, "/") > 0 {
		req_url = "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/root:" + path + ":/children"
	} else {
		req_url = "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/items/" + path + "/children"
	}
	r, err := req.Get(req_url, header)
	if err != nil {
		log.Println(err)
	}
	var response models.ChildResponse
	_ = json.Unmarshal(r.Bytes(), &response)
	return response.Value
}

func ApiGetDriveItem(c *gin.Context) {
	var request struct{
		Path string `json:"path"`
	}
	_ = c.BindJSON(&request)

	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    GetDriveItem(request.Path),
	})
}

func ApiPutDriveItem(c *gin.Context) {
	var request struct{
		ItemID string `json:"item_id"`
	}
	_ = c.BindJSON(&request)
	viper.Set("ItemID", request.ItemID)
	viper.WriteConfig()
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "Item ID set.",
		Data:    nil,
	})
}

func GetUserInfo() (models.UserInfo) {
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	r, err := req.Get("https://graph.microsoft.com/v1.0/me", header)
	if err != nil {
		log.Println(err)
	}
	var user_info models.UserInfo
	r.ToJSON(&user_info)
	return user_info
}