package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

type MeetingJoinService struct {
	MeetingId string `form:"meetingid" json:"meetingid" binding:"required"`
	UserId    string `form:"userid" json:"userid" binding:"required"`
}

func (service *MeetingJoinService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到user", nil)
		return &rtn
	}

	return nil

}

func (service *MeetingJoinService) JoinMeeting(c *gin.Context) serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var meeting model.Meeting

	// 查询meeting
	if err := model.DB.Where("mid = ?", service.MeetingId).First(&meeting).Error; err != nil {
		return serializer.ParamErr("Meeting Not Found", nil)
	}

	// check if exist
	count := 0
	model.DB.Model(&model.Relation{}).Where("user_id = ? AND meeting_id = ?", service.UserId, service.MeetingId).Count(&count)
	if count > 0 {
		return serializer.ParamErr("user already in this meeting !", nil)
	}

	// add member to memberlist
	relation := model.Relation{
		UserId:    service.UserId,
		MeetingId: service.MeetingId,
		Type:      0,
	}
	if err := model.DB.Save(&relation).Error; err != nil {
		return serializer.Err(40002, "保存失败", err)
	}

	detailService := MeetingDetialService{
		Meetingid: service.MeetingId,
	}
	return detailService.GetDetail(c)
}
