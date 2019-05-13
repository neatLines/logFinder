package models

//
type Filter struct {
	Hostname  []string `json:"hostname"`
	StartTime string   `json:"startTime"`
	EndTime   string   `json:"endTime"`
	Filter    string   `json:"filter"`
	NeedFlush bool     `json:"needFlush"`
	AppName   string   `json:"appName"`
}
