package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"strings"
)

type MeetingLeaveService struct{
	MeetingId string `form:"meetingid" json:"meetingid" binding:"required"`
	UserId string `form:"userid" json:"userid" binding:"required"`
}

func (service *MeetingLeaveService) valid() *serializer.Response{
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到user", nil)
		return &rtn
	}

	return nil

}

func (service *MeetingLeaveService) LeaveMeeting(c *gin.Context) serializer.Response{
	if err := service.valid(); err != nil{
		return *err
	}

	var meeting model.Meeting

	// 查询meeting
	if err:= model.DB.Where("mid = ?", service.MeetingId).First(&meeting).Error; err!= nil{
		return serializer.ParamErr("Meeting Not Found",nil)
	}

	// build new member list
	members := strings.Split(meeting.MemberList,",")
	var newMemberList string
	inMemberList := false
	for _,v := range members {
		if v == service.UserId {
			inMemberList = true
		} else{
			newMemberList += "," + v
		}
	}

	if !inMemberList {
		return serializer.Err(40003, "用户不在该meeting的memberlist中",nil)
	}

	newMemberList = newMemberList[1:]
	meeting.MemberList = newMemberList

	if err := model.DB.Save(&meeting).Error; err != nil{
		return serializer.Err(40002, "保存失败", err)
	}

	detailService := MeetingDetialService{
		Meetingid:service.MeetingId,
	}
	return detailService.GetDetail(c)
}
