package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MeetingUpdateService struct {
	gorm.Model
	MeetingId string `form:"mid" json:"mid" binding:"required"`
	MeetingName string `form:"mname" json:"mname" binding:"required"`
	MeetingCover string `form:"mcover" json:"mcover" binding:"required"`
	CheckType string `form:"check_type" json:"check_type" binding:"required"`
	CheckRule string `form:"check_rule" json:"check_rule" binding:"required"`
	TimeStart string `form:"check_time_start" json:"check_time_start" binding:"required"`
	TimeEnd string `form:"check_time_end" json:"check_time_end" binding:"required"`
	Longitude string `form:"longitude" json:"longitude" binding:"required"`
	Latitude string `form:"latitude" json:"latitude" binding:"required"`
	Describe string `form:"describe" json:"describe" binding:"required"`
}

func (service MeetingUpdateService) UpdateMeeting(c *gin.Context) serializer.Response {
	// fetch old meeting by mid
	var meeting model.Meeting
	if err := model.DB.Where("mid = ?", service.MeetingId).First(&meeting).Error; err != nil{
		return serializer.ParamErr("未找到相应meeting", err)
	}

	// update meeting
	meeting.MeetingName = service.MeetingName
	meeting.MeetingCover = service.MeetingCover
	meeting.CheckType = service.CheckType
	meeting.CheckRule = service.CheckRule
	meeting.TimeStart = service.TimeStart
	meeting.TimeEnd = service.TimeEnd
	meeting.LocationLongitude = service.Longitude
	meeting.LocationLatitude = service.Latitude
	meeting.Describe = service.Describe

	model.DB.Save(&meeting)
	detailService := MeetingDetialService{
		Meetingid: service.MeetingId,
	}
	return detailService.GetDetail(c)
}

