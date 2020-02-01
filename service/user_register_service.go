package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"fmt"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	PhoneNumber  string `form:"phonenumber" json:"phonenumber" binding:"required"`
	UserName     string `form:"username" json:"username" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Face         string `form:"face" json:"face" binding:"required"`
	Sex          uint8  `form:"sex" json:"sex" binding:"required"`
	Organization string `form:"organization" json:"organization" binding:"required"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.PhoneNumber).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "手机号已注册",
		}
	}

	if service.Sex != 1 && service.Sex != 2 {
		return &serializer.Response{
			Code: 40002,
			Msg:  "性别只能为1或2",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		PhoneNumber:  service.PhoneNumber,
		UserName:     service.UserName,
		Face:         service.Face,
		Sex:          service.Sex,
		Organization: service.Organization,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}
	fmt.Println("------------")
	fmt.Println(service.Password)
	fmt.Println("------------")

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
