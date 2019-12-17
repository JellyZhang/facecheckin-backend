package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

type CheckAddService struct {
	UserId       string `form:"uid" json:"uid" binding:"required"`
	MeetingId    string `form:"mid" json:"mid" binding:"required"`
	CheckType    int `form:"type" json:"type" binding:"required"`
}
// valid 验证表单
func (service *CheckAddService) valid() *serializer.Response {
	count := 0
	if model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count); count == 0 {
		rtn := serializer.ParamErr("用户未找到",nil)
		return &rtn
	}
	count = 0
	if model.DB.Model(&model.Meeting{}).Where("mid = ?", service.MeetingId).Count(&count); count == 0 {
		rtn := serializer.ParamErr("meeting未找到",nil)
		return &rtn
	}
	if service.CheckType != 2 && service.CheckType != 1 {
		rtn := serializer.ParamErr("check type 必须为0或1",nil)
		return &rtn
	}


	// 	签退时检查是否已签到
	if service.CheckType == 2 {
		count = 0
		model.DB.Model(&model.Check{}).Where("created_at >= ? AND created_at < ? AND check_type = 1", time.Now().Format("2006-01-02"), time.Now()).Count(&count)
		if count == 0 {
			rtn := serializer.ParamErr("在签退前必须先签到", nil)
			return &rtn
		}
	}
	count = 0
	model.DB.Model(&model.Check{}).Where("created_at >= ? AND created_at < ? AND check_type = 2", time.Now().Format("2006-01-02"), time.Now()).Count(&count)
	if count > 0 {
		rtn := serializer.ParamErr("今天已签退", nil)
		return &rtn
	}

	return nil
}

func (service *CheckAddService) AddCheck(c *gin.Context) serializer.Response {

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	//const TIME_LAYOUT = "2006-01-02 15:04:05"
	//newchecktime, err := time.Parse(TIME_LAYOUT,service.CheckTime)
	//if err != nil{
	//	return serializer.ParamErr("时间转换失败", err)
	//}
	newcheck := model.Check{
		UserId:    service.UserId,
		MeetingId: service.MeetingId,
		//CheckTime: newchecktime,
		CheckType: service.CheckType,
	}
	var whichtype string
	if service.CheckType == 1 {
		whichtype = "签到"
	} else{
		whichtype = "签退"
	}

	if ok := model.DB.NewRecord(newcheck); ok {
		if err := model.DB.Create(&newcheck).Error; err != nil {
			return serializer.ParamErr("添加"+whichtype + "失败1", err)
		}
	} else {
		return serializer.ParamErr("添加"+whichtype+"失败 : 已存在记录", nil)
	}

	return serializer.ParamGood("添加"+whichtype+"成功")

}