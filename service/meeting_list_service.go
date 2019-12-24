package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

type MeetingListService struct {
	UserId string `form:"uid" json:"uid" binding:"required"`
	Type string `form:"type" json:"type" binding:"required"`
}
func (service *MeetingListService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到user", nil)
		return &rtn
	}

	return nil

}
func (service MeetingListService) GetList (c *gin.Context) serializer.Response{
	if err := service.valid(); err != nil {
		return *err
	}

	var relations []model.Relation
	var meetings []model.Meeting
	if err:= model.DB.Where("user_id = ? && type = ?",service.UserId, service.Type).Find(&relations).Error; err!=nil{
		return serializer.Err(50000,"查找失败",err)
	}
	for _,r := range relations{
		var temp model.Meeting
		if err:= model.DB.Where("mid = ?",r.MeetingId).First(&temp).Error; err != nil{
			return serializer.Err(50000,"查找失败",err)
		}
		meetings = append(meetings, temp)
	}
	return serializer.BuildMeetingListResponse(meetings)
}