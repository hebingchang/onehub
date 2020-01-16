package models

type OkJson struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ErrorJson struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
}
