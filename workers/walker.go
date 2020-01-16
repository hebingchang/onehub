package workers

import (
	"encoding/json"
	"github.com/imroc/req"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"onehub/models"
	"time"
)

func token() (models.Token) {
	var token models.Token
	_ = mapstructure.Decode(viper.Get("Token"), &token)
	return token
}

func WalkDrive() (models.Folder) {
	return walk("存储", time.Time{}, viper.GetString("ItemID"), "")
}

func walk(name string, update_time time.Time, id string, path string) models.Folder {
	log.Printf("walking %s (%s)", name, id)
	if path == "" { path = "/" } else {
		if path == "/" {
			path = path + name
		} else {
			path = path + "/" + name
		}
	}
	fs := models.Folder{
		Name:       name,
		UpdateTime: update_time,
		ID:         id,
		Path:       path,
	}
	var files []models.File
	var folders []models.Folder

	header := req.Header{
		"Authorization": "Bearer " + token().AccessToken,
	}
	r, err := req.Get("https://graph.microsoft.com/v1.0/drives/" + viper.GetString("DriveID") + "/items/" + id +  "/children", header)
	if err != nil {
		log.Println(err)
	}
	var drive_file models.DriveFileResponse
	_ = json.Unmarshal(r.Bytes(), &drive_file)

	for _, item := range drive_file.Value {
		if item.MicrosoftGraphDownloadURL == "" {
			// is a directory
			folders = append(folders, walk(item.Name, item.LastModifiedDateTime, item.ID, path))
		} else {
			// is a file
			files = append(files, models.File{
				Name:       item.Name,
				UpdateTime: item.LastModifiedDateTime,
				ID:         item.ID,
				Size:       item.Size,
			})
		}
	}
	fs.Files = files
	fs.Folders = folders
	return fs
}
