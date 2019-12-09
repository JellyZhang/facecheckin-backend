package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
)

// UserUpdateService 管理用户信息更新服务
type UserUpdateService struct {
	PhoneNumber  string `form:"phonenumber" json:"phonenumber" binding:"required,min=2,max=30"`
	UserName     string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password     string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Face         string `form:"face" json:"face" binding:"required"`
	Sex          uint8  `form:"sex" json:"sex" binding:"required"`
	Organization string `form:"organization" json:"organization" binding:"required"`
}

// valid 验证表单
func (service *UserUpdateService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ? AND phone_number != ?", service.UserName, service.PhoneNumber).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	if service.Sex != 1 && service.Sex != 2 {
		return &serializer.Response{
			Code:  40002,
			Msg:   "性别只能为1或2",
		}
	}

	return nil
}

// Update 用户信息更新
func (service *UserUpdateService) Update(c *gin.Context) serializer.Response {
	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 查找用户
	var user model.User
	if err := model.DB.Where("phone_number = ?", service.PhoneNumber).First(&user).Error; err != nil {
		return serializer.ParamErr("未找到用户",err)
	}

	// 更新信息
	user.UserName = service.UserName
	user.Organization = service.Organization
	user.Sex = service.Sex
	user.Face = service.Face
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 保存信息
	if err := model.DB.Save(&user).Error; err != nil{
		return serializer.ParamErr("保存信息失败",err)
	}

	return serializer.BuildUserResponse(user)
}
