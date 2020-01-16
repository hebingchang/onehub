package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"onehub/config"
	"onehub/models"
	"time"
)

func getAuthUrl(redirect_uri string) (string) {
	oauth_url := url.URL{
		Scheme:     "https",
		Host:       "login.microsoftonline.com",
		Path:       "/common/oauth2/v2.0/authorize",
	}
	q := oauth_url.Query()
	q.Set("client_id", config.CLIENT_ID)
	q.Set("response_type", "code")
	q.Set("redirect_uri", redirect_uri)
	q.Set("response_mode", "query")
	q.Set("scope", "offline_access user.read files.read.all files.readwrite.all")
	q.Set("state", "state_at_ease")
	oauth_url.RawQuery = q.Encode()
	return oauth_url.String()
}

func OauthRedirect(c *gin.Context) {
	u := location.Get(c)
	c.Redirect(301, getAuthUrl(fmt.Sprintf("%s://%s/api/v1/admin/oauth/callback", u.Scheme, u.Host)))
}

func OauthCallback(c *gin.Context) {
	u := location.Get(c)
	param := req.Param{
		"client_id": config.CLIENT_ID,
		"scope":  "offline_access user.read files.read.all files.readwrite.all",
		"code":  c.Query("code"),
		"redirect_uri":  fmt.Sprintf("%s://%s/api/v1/admin/oauth/callback", u.Scheme, u.Host),
		"grant_type":  "authorization_code",
		"client_secret":  config.CLIENT_SECRET,
	}

	// http://localhost:8080/api/v1/admin/oauth
	var token models.Token
	response, err := req.Post("https://login.microsoftonline.com/common/oauth2/v2.0/token", param)
	if err != nil {
		log.Fatal(err)
	}
	_ = response.ToJSON(&token)
	token.ExpiresAt = int(time.Now().Unix()) + token.ExpiresIn
	viper.Set("Token", &token)
	viper.WriteConfig()
	user_info, _ := json.Marshal(GetUserInfo())

	c.HTML(200, "callback.tmpl", gin.H{
		"user_info": string(user_info),
		"host": fmt.Sprintf("%s://%s", u.Scheme, u.Host),
	})
}

func OauthProfile(c *gin.Context) {
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    GetUserInfo(),
	})
}

func GetCurrentDrive(c *gin.Context) {
	c.JSON(200, models.OkJson{
		Success: true,
		Message: "ok",
		Data:    viper.GetString("DriveID"),
	})
}

func GetCurrentItem(c *gin.Context) {
	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	req_url := "https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/items/" + viper.GetString("ItemID")
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