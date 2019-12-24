package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

type MeetingDetialService struct {
	Meetingid string `form:"meetingid" json:"meetingid" binding:"required"`
}

func (service *MeetingDetialService) GetDetail(c *gin.Context) serializer.Response {
	var meeting model.Meeting
	var userList []model.User

	// 查询meeting
	if err := model.DB.Where("mid = ?", service.Meetingid).First(&meeting).Error; err != nil {
		return serializer.ParamErr("Meeting Not Found", err)
	}

	// find member list
	var relations []model.Relation
	if err := model.DB.Where("meeting_id = ? AND type = 0", meeting.Mid).Find(&relations).Error; err != nil {
		return serializer.ParamErr("cant find member list", err)
	}

	for _, relation := range relations {
		var user model.User
		if err := model.DB.Where("phone_number = ?", relation.UserId).First(&user).Error; err == nil {
			userList = append(userList, user)
		}
	}

	// find owner
	var owner model.User
	if err := model.DB.Where("phone_number = ?", meeting.OwnerId).First(&owner).Error; err != nil {
		return serializer.ParamErr("cant find owner", err)
	}

	// build meetingDetail
	detail := model.MeetingDetail{
		MeetingCover: meeting.MeetingCover,
		MeetingName:  meeting.MeetingName,
		Mid:          meeting.Mid,
		MeetingLocation: model.MeetingLocation{
			Longitude: meeting.LocationLongitude,
			Latitude:  meeting.LocationLatitude,
			Describe:  meeting.Describe,
		},
		CheckTime: model.Checktime{
			CheckRule: meeting.CheckRule,
			TimeSpace: model.TimeSpace{
				TimeStart: meeting.TimeStart,
				TimeEnd:   meeting.TimeEnd,
			},
		},
		Owner:      owner,
		MemberList: userList,
	}

	return serializer.BuildMeetingDetailResponse(detail)
}
