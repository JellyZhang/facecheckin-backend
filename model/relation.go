package model

import "github.com/jinzhu/gorm"

type Relation struct {
	gorm.Model
	UserId string
	MeetingId   string
	Type        int
}
