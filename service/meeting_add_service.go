package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type MeetingAddService struct {
	UserId       string `form:"uid" json:"uid" binding:"required"`
	MeetingName  string `form:"mname" json:"mname" binding:"required"`
	MeetingCover string `form:"mcover" json:"mcover" binding:"required"`
	CheckRule    string `form:"check_rule" json:"check_rule" binding:"required"`
	TimeStart    int `form:"check_time_start" json:"check_time_start" binding:"required"`
	TimeEnd      int `form:"check_time_end" json:"check_time_end" binding:"required"`
	Longitude    string `form:"longitude" json:"longitude" binding:"required"`
	Latitude     string `form:"latitude" json:"latitude" binding:"required"`
	Describe     string `form:"describe" json:"describe" binding:"required"`
}

const lenOfMid = 6

//  get available meeting id
func getAvailaleMid() string {
	rand.Seed(time.Now().Unix())
	limit := int(math.Pow10(lenOfMid))
	//var meeting model.Meeting
	for i := 1; i < limit; {
		temp := strconv.Itoa(rand.Intn(limit))
		for len(temp) < lenOfMid {
			temp = "0" + temp
		}
		count := 0
		model.DB.Model(&model.Meeting{}).Where("mid = ?", temp).Count(&count)
		if count == 0 {
			return temp
		}
	}
	return "available mid not found"
}

// valid 验证表单
func (service *MeetingAddService) valid() *serializer.Response {
	if service.TimeStart < 0 || service.TimeStart > 1440 || service.TimeEnd < 0 || service.TimeEnd > 1440 || service.TimeStart > service.TimeEnd {
		rtn := serializer.ParamErr("时间范围不符合规范", nil)
		return &rtn
	}
	var owner model.User
	if err := model.DB.Where("phone_number = ?", service.UserId).First(&owner).Error; err != nil {
		rtn := serializer.ParamErr("主持人未找到", err)
		return &rtn
	}
	return nil
}

func (service MeetingAddService) AddMeeting(c *gin.Context) serializer.Response {

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 寻找可用Mid
	fmt.Println("寻找可用mid")
	newMid := getAvailaleMid()
	if len(newMid) > lenOfMid {
		return serializer.ParamErr("可用的Mid未找到", nil)
	}
	fmt.Println(newMid)

	// build meeting
	newMeeting := model.Meeting{
		Model:             gorm.Model{},
		MeetingCover:      service.MeetingCover,
		MeetingName:       service.MeetingName,
		Mid:               newMid,
		LocationLatitude:  service.Latitude,
		LocationLongitude: service.Longitude,
		Describe:          service.Describe,
		CheckRule:         service.CheckRule,
		TimeStart:         service.TimeStart,
		TimeEnd:           service.TimeEnd,
		OwnerId:             service.UserId,
	}
	newrelation := model.Relation{
		UserId:    service.UserId,
		MeetingId: newMid,
		Type:      1,
	}
	if ok := model.DB.NewRecord(newMeeting); ok {
		if err := model.DB.Create(&newMeeting).Error; err != nil {
			return serializer.ParamErr("添加会议失败1", err)
		}
	} else {
		return serializer.ParamErr("添加会议失败 : 已存在记录", nil)
	}
	if ok := model.DB.NewRecord(newrelation); ok {
		if err := model.DB.Create(&newrelation).Error; err != nil {
			return serializer.ParamErr("添加owner失败1", err)
		}
	} else {
		return serializer.ParamErr("添加owner失败 : 已存在记录", nil)
	}
	return serializer.ParamGood("添加会议成功, meeting id = " + newMid)
}
