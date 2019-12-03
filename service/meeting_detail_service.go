package service

import (
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MeetingDetialService struct{
	gorm.Model
	Meetingid string `form:"meetingid" json:"meetingid" binding:"required"`
}

func (service *MeetingDetialService) GetDetail(c *gin.Context) serializer.Response{
	return serializer.BuildMeetingDetailResponse(service.Meetingid)
	//if err := model.DB.Where("phone_number = ?", service.PhoneNumber).First(&user).Error; err != nil {
	//	return serializer.ParamErr("手机号或密码错误1", nil)
	//}
	//
	//if user.CheckPassword(service.Password) == false {
	//	return serializer.ParamErr("手机号或密码错误2", nil)
	//}


	//return serializer.BuildUserResponse(user)

}
