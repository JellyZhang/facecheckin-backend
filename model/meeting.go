package model

import "github.com/jinzhu/gorm"

type Meeting struct {
	gorm.Model
	MeetingCover      string
	MeetingName       string
	Mid               string
	LocationLatitude  string
	LocationLongitude string
	Describe          string
	CheckRule         string
	TimeStart         int
	TimeEnd           int
	OwnerId           string
}
