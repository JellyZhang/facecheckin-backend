package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type MeetingDetialService struct{
	gorm.Model
	Meetingid string `form:"meetingid" json:"meetingid" binding:"required"`
}

func (service *MeetingDetialService) GetDetail(c *gin.Context) serializer.Response{
	var meeting model.Meeting
	var userList []model.User

	// 查询meeting
	if err:= model.DB.Where("mid = ?", service.Meetingid).First(&meeting).Error; err!= nil{
		return serializer.ParamErr("Meeting Not Found",nil)
	}

	// 将memberlist的1,2,3形式切割并转化为[]user
	members := strings.Split(meeting.MemberList,",")
	for _, member := range members{
		var tempUser model.User
		if uid, err := strconv.Atoi(member); err == nil {
			if err := model.DB.Where("id = ?", uid).First(&member).Error; err!=nil{
				userList = append(userList, tempUser)
			}
		}
	}

	// build meetingDetail
	detail := model.MeetingDetail{
		MeetingCover:    meeting.MeetingCover,
		MeetingName:     meeting.MeetingName,
		Mid:             meeting.Mid,
		MeetingLocation: model.MeetingLocation{
			Longitude: meeting.LocationLongitude,
			Latitude:  meeting.LocationLatitude,
			Describe:  meeting.Describe,
		},
		CheckTime:       model.Checktime{
			CheckType: meeting.CheckType,
			CheckRule: meeting.CheckRule,
			TimeSpace: model.TimeSpace{
				TimeStart: meeting.TimeStart,
				TimeEnd:   meeting.TimeEnd,
			},
		},
		MemberList:      userList,
	}

	return serializer.BuildMeetingDetailResponse(detail)
}
