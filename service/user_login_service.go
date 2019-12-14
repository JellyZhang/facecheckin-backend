package service

import (
	"facecheckin/model"
	"facecheckin/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("phone_number = ?", service.PhoneNumber).First(&user).Error; err != nil {
		return serializer.ParamErr("手机号或密码错误1", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("手机号或密码错误2", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
