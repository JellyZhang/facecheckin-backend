package model

type Checktime struct {
	CheckType string      `json:"checktype"`
	CheckRule string      `json:"checkrule"`
	TimeSpace interface{} `json:"timespace"`
}

type TimeSpace struct {
	TimeStart int `json:"start"`
	TimeEnd   int `json:"end"`
}
type MeetingLocation struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Describe  string `json:"describe"`
}
type MeetingDetail struct {
	MeetingCover    string      `json:"mcover"`
	MeetingName     string      `json:"mname"`
	Mid             string      `json:"mid"`
	MeetingLocation interface{} `json:"location"`
	CheckTime       interface{} `json:"checktime"`
	MemberList      []User
	Owner           User
}
