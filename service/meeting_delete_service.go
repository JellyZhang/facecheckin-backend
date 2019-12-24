package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

type MeetingDeleteService struct {
	MeetingId string `form:"meetingid" json:"meetingid" binding:"required"`
	UserId    string `form:"userid" json:"userid" binding:"required"`
}
func (service *MeetingDeleteService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到user", nil)
		return &rtn
	}
	var meeting model.Meeting
	if err := model.DB.Where("mid = ?", service.MeetingId).First(&meeting).Error; err != nil {
		rtn := serializer.ParamErr("未找到相应meeting", err)
		return &rtn
	}

	if meeting.OwnerId != service.UserId {
		rtn := serializer.Err(40001, "您不是该会议的主持人", nil)
		return &rtn
	}
	return nil

}
func (service MeetingDeleteService) Delete (c *gin.Context) serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	if err := model.DB.Where("meeting_id = ?",service.MeetingId).Delete(&model.Relation{}).Error; err != nil{
		return serializer.Err(50000, "删除relation失败", err)
	}
	if err := model.DB.Where("meeting_id = ?",service.MeetingId).Delete(&model.Check{}).Error; err != nil{
		return serializer.Err(50000, "删除check失败", err)
	}
	if err := model.DB.Where("mid = ?",service.MeetingId).Delete(&model.Meeting{}).Error; err != nil{
		return serializer.Err(50000, "删除meeting失败", err)
	}

	return serializer.ParamGood("删除成功")

}