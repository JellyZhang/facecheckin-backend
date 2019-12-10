package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"strings"
)

type MeetingJoinService struct{
	MeetingId string `form:"meetingid" json:"meetingid" binding:"required"`
	UserId string `form:"userid" json:"userid" binding:"required"`
}

func (service *MeetingJoinService) valid() *serializer.Response{
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到user", nil)
		return &rtn
	}

	return nil

}

func (service *MeetingJoinService) JoinMeeting(c *gin.Context) serializer.Response{
	if err := service.valid(); err != nil{
		return *err
	}

	var meeting model.Meeting

	// 查询meeting
	if err:= model.DB.Where("mid = ?", service.MeetingId).First(&meeting).Error; err!= nil{
		return serializer.ParamErr("Meeting Not Found",nil)
	}

	// check if exist
	members := strings.Split(meeting.MemberList,",")
	for _,v := range members {
		if v == service.UserId {
			return serializer.Err(40003, "用户已加入该meeting", nil)
		}
	}

	// add member to memberlist
	if len(meeting.MemberList) == 0 {
		meeting.MemberList = service.UserId
	} else{
		meeting.MemberList += "," + service.UserId
	}
	if err := model.DB.Save(&meeting).Error; err != nil{
		return serializer.Err(40002, "保存失败", err)
	}

	detailService := MeetingDetialService{
		Meetingid:service.MeetingId,
	}
	return detailService.GetDetail(c)
}
