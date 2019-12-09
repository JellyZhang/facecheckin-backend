package model

import "github.com/jinzhu/gorm"

type Meeting struct{
	gorm.Model
	MeetingCover string
	MeetingName string
	Mid string
	LocationLatitude string
	LocationLongitude string
	Describe string
	CheckType string
	CheckRule string
	TimeStart string
	TimeEnd string
	Owner string
	MemberList string
}
