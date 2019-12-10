package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type MeetingDetialService struct{
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
	fmt.Println("memberlist = " + meeting.MemberList)
	members := strings.Split(meeting.MemberList,",")
	for _,v := range members {
		fmt.Println(v)
	}
	members =  append(members[0:1],members[0:]...)
	members[0] = meeting.Owner
	for _, member := range members{
		var tempUser model.User
		if err := model.DB.Where("phone_number = ?", member).First(&tempUser).Error; err ==nil{
			userList = append(userList, tempUser)
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
