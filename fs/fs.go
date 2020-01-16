package fs

import (
	"encoding/gob"
	"errors"
	"fmt"
	"onehub/models"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Root models.Folder
var LastUpdate time.Time

type DataFile struct {
	Fs           models.Folder
	LastUpdate   time.Time
}

func init() {
	dataFile, err := os.OpenFile("fs.gob", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data_file DataFile
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data_file)

	if err != nil {
		fmt.Println(err)
	}

	Root = data_file.Fs
	LastUpdate = data_file.LastUpdate

	_ = dataFile.Close()
}

func WriteFsDataFile() {
	dataFile, err := os.OpenFile("fs.gob", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataEncoder := gob.NewEncoder(dataFile)
	_ = dataEncoder.Encode(DataFile{
		Fs:         Root,
		LastUpdate: LastUpdate,
	})

	_ = dataFile.Close()
}

type children struct {
	Files []models.File `json:"files"`
	Folders []models.Folder `json:"folders"`
}

func GetChildren(path string) (*children, error) {
	path, _ = filepath.Abs(path)
	folder := &Root
	if len(strings.Split(path, "/")) == 1 || strings.Split(path, "/")[1] == "" {
		return &children{
			Files:   folder.Files,
			Folders: folder.Folders,
		}, nil
	}
	for _, name := range strings.Split(path, "/")[1:] {
		if name != "" {
			found := false
			for _, item := range folder.Folders {
				if item.Name == name {
					folder = &item
					found = true
					break
				}
			}
			if !found {
				return nil, errors.New("未找到目录")
			}
		}
	}
	return &children{
		Files:   folder.Files,
		Folders: folder.Folders,
	}, nil
}