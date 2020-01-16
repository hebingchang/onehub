package models

import (
	"onehub/utils"
	"time"
)

type UserInfo struct {
	OdataContext      string        `json:"@odata.context"`
	BusinessPhones    []interface{} `json:"businessPhones"`
	DisplayName       string        `json:"displayName"`
	GivenName         string        `json:"givenName"`
	JobTitle          interface{}   `json:"jobTitle"`
	Mail              string        `json:"mail"`
	MobilePhone       interface{}   `json:"mobilePhone"`
	OfficeLocation    interface{}   `json:"officeLocation"`
	PreferredLanguage interface{}   `json:"preferredLanguage"`
	Surname           string        `json:"surname"`
	UserPrincipalName string        `json:"userPrincipalName"`
	ID                string        `json:"id"`
}

type Token struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    int    `json:"expires_at"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Drive struct {
	CreatedDateTime      time.Time `json:"createdDateTime"`
	Description          string    `json:"description"`
	ID                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Name                 string    `json:"name"`
	WebURL               string    `json:"webUrl"`
	DriveType            string    `json:"driveType"`
	CreatedBy            struct {
		User struct {
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	Owner struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"owner"`
	Quota struct {
		Deleted   int64  `json:"deleted"`
		Remaining int64  `json:"remaining"`
		State     string `json:"state"`
		Total     int64  `json:"total"`
		Used      int64  `json:"used"`
	} `json:"quota"`
}

type Child struct {
	CreatedDateTime      time.Time `json:"createdDateTime"`
	ETag                 string    `json:"eTag"`
	ID                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Name                 string    `json:"name"`
	WebURL               string    `json:"webUrl"`
	CTag                 string    `json:"cTag"`
	Size                 int64     `json:"size"`
	CreatedBy            struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
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
	FileSystemInfo struct {
		CreatedDateTime      time.Time `json:"createdDateTime"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"fileSystemInfo"`
	Folder struct {
		ChildCount int `json:"childCount"`
	} `json:"folder"`
}

type ChildResponse struct {
	OdataContext string        `json:"@odata.context"`
	Value        []interface{} `json:"value"`
}

type DriveResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []Drive `json:"value"`
}

type DriveItem struct {
	OdataContext         string    `json:"@odata.context"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	ETag                 string    `json:"eTag"`
	ID                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Name                 string    `json:"name"`
	WebURL               string    `json:"webUrl"`
	CTag                 string    `json:"cTag"`
	Size                 int64     `json:"size"`
	CreatedBy            struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	Folder struct {
		ChildCount int `json:"childCount"`
	} `json:"folder"`
	MicrosoftGraphDownloadURL string    `json:"@microsoft.graph.downloadUrl"`
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
	FileSystemInfo struct {
		CreatedDateTime      time.Time `json:"createdDateTime"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"fileSystemInfo"`
	Video struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"video"`
}

type Config struct {
	RefreshInterval int    `json:"refresh_interval"`
	DriveID         string `json:"drive_id"`
	ItemID          string `json:"item_id"`
}

type SessionItem struct {
	OdataType                      string `json:"@odata.type"`
	MicrosoftGraphConflictBehavior string `json:"@microsoft.graph.conflictBehavior"`
	Name                           string `json:"name"`
}

type CreateSessionResponse struct {
	ExpirationDateTime time.Time `json:"expirationDateTime"`
	NextExpectedRanges []string  `json:"nextExpectedRanges"`
	UploadURL          string    `json:"uploadUrl"`
}

type CreateSessionRequest struct {
	Item SessionItem `json:"item"`
}

func (d Drive) String() string {
	return d.Name + " (" + d.DriveType + ", " + utils.QuotaFormat(d.Quota.Used) + "/" + utils.QuotaFormat(d.Quota.Total) + ") " + d.WebURL
}