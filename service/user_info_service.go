package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

// UserInfoService 管理用户登录的服务
type UserInfoService struct {
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
}

// UserInfo 用户信息
func (service *UserInfoService) Info(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("phone_number = ?", service.PhoneNumber).First(&user).Error; err != nil {
		return serializer.ParamErr("手机号未注册", nil)
	}

	return serializer.BuildUserResponse(user)
}
