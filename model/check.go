package model

import (
	"github.com/jinzhu/gorm"
)

// 签到记录模型
type Check struct {
	gorm.Model
	UserId string
	MeetingId string
	CheckType int
}