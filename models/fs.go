package models

import "time"

type DriveFileResponse struct {
	OdataContext string      `json:"@odata.context"`
	Value        []DriveFile `json:"value"`
}

type DriveFile struct {
	FileSystemInfo struct {
		CreatedDateTime      time.Time `json:"createdDateTime"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"fileSystemInfo"`
	Folder struct {
		ChildCount int `json:"childCount"`
	} `json:"folder"`
	MicrosoftGraphDownloadURL string    `json:"@microsoft.graph.downloadUrl"`
	CreatedDateTime           time.Time `json:"createdDateTime"`
	ETag                      string    `json:"eTag"`
	ID                        string    `json:"id"`
	LastModifiedDateTime      time.Time `json:"lastModifiedDateTime"`
	Name                      string    `json:"name"`
	WebURL                    string    `json:"webUrl"`
	CTag                      string    `json:"cTag"`
	Size                      int64     `json:"size"`
	CreatedBy                 struct {
		Application struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"application"`
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		Application struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"application"`
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	ParentReference struct {
		DriveID   string `json:"driveId"`
		DriveType string `json:"driveType"`
		ID        string `json:"id"`
		Path      string `json:"path"`
	} `json:"parentReference"`
	File struct {
		MimeType string `json:"mimeType"`
		Hashes   struct {
			QuickXorHash string `json:"quickXorHash"`
		} `json:"hashes"`
	} `json:"file"`
	Shared struct {
		Scope string `json:"scope"`
	} `json:"shared"`
}

type File struct {
	Name        string `json:"name"`
	UpdateTime  time.Time `json:"update_time"`
	ID          string `json:"id"`
	Size        int64  `json:"size"`
}

type Folder struct {
	Name    string   `json:"name"`
	UpdateTime  time.Time `json:"update_time"`
	Files   []File   `json:"-"`
	Folders []Folder `json:"-"`
	ID      string   `json:"id"`
	Path    string   `json:"path"`
}

type FolderTree struct {
	Name    string   `json:"name"`
	Folders []FolderTree `json:"folders,omitempty"`
	ID      string   `json:"id"`
}

func (f Folder) Tree(pathArray ...string) FolderTree {
	var path string
	if len(pathArray) == 0 {
		path = ""
	} else {
		path = pathArray[0]
	}
	tree := FolderTree{
		Name:    f.Name,
		ID:      path,
	}
	var folders []FolderTree
	if len(f.Folders) > 0 {
		for _, folder := range(f.Folders) {
			folders = append(folders, folder.Tree(path + "/" + folder.Name))
		}
	}
	tree.Folders = folders
	return tree
}

func (f Folder) FoldersCount() int {
	count := len(f.Folders)
	if count != 0 {
		for _, folder := range(f.Folders) {
			count += folder.FoldersCount()
		}
		return count
	} else {
		return 0
	}
}

func (f Folder) FilesCount() int {
	count := len(f.Files)
	if len(f.Folders) != 0 {
		for _, folder := range(f.Folders) {
			count += folder.FilesCount()
		}
		return count
	} else {
		return count
	}
}