package services

import (
	"github.com/imroc/req"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"onehub/models"
	"time"
)

func RefreshToken() {
	var token models.Token
	_ = mapstructure.Decode(viper.Get("Token"), &token)

	param := req.Param{
		"client_id": viper.GetString("ClientID"),
		"scope":  "offline_access user.read files.read.all files.readwrite.all",
		"refresh_token": token.RefreshToken,
		"grant_type":  "refresh_token",
		"client_secret":  viper.GetString("ClientSecret"),
	}

	response, err := req.Post("https://login.microsoftonline.com/common/oauth2/v2.0/token", param)
	if err != nil {
		log.Fatal(err)
	}
	_ = response.ToJSON(&token)
	token.ExpiresAt = int(time.Now().Unix()) + token.ExpiresIn
	viper.Set("Token", &token)
	viper.WriteConfig()
}

