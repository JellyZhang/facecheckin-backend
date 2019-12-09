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
)

type MeetingAddService struct {
	gorm.Model
	UserId string `form:"uid" json:"uid" binding:"required"`
	MeetingName string `form:"mname" json:"mname" binding:"required"`
	MeetingCover string `form:"mcover" json:"mcover" binding:"required"`
	CheckType string `form:"check_type" json:"check_type" binding:"required"`
	CheckRule string `form:"check_rule" json:"check_rule" binding:"required"`
	TimeStart string `form:"check_time_start" json:"check_time_start" binding:"required"`
	TimeEnd string `form:"check_time_end" json:"check_time_end" binding:"required"`
	Longitude string `form:"longitude" json:"longitude" binding:"required"`
	Latitude string `form:"latitude" json:"latitude" binding:"required"`
	Describe string `form:"describe" json:"describe" binding:"required"`
}
const lenOfMid = 6

func getAvailaleMid() string{
	rand.Seed(42)
	limit := int(math.Pow10(lenOfMid))
	//var meeting model.Meeting
	for i:=1; i < limit
	{
		temp := strconv.Itoa(rand.Intn(limit))
		for len(temp) < lenOfMid {
			temp = "0" + temp
		}
		if err := model.DB.Where("mid = ?", temp).Error; err == nil{
			return temp
		}
	}
	return "available mid not found"
}
func (service MeetingAddService) AddMeeting(c *gin.Context) serializer.Response {
	// 寻找可用Mid
	fmt.Println("寻找可用mid")
	newMid := getAvailaleMid()
	if len(newMid) > lenOfMid {
		return serializer.ParamErr("可用的Mid未找到",nil)
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
		CheckType:         service.CheckType,
		CheckRule:         service.CheckRule,
		TimeStart:         service.TimeStart,
		TimeEnd:           service.TimeEnd,
		Owner:             service.UserId,
		MemberList:        "",
	}
	if ok := model.DB.NewRecord(newMeeting); ok {
		if err := model.DB.Create(&newMeeting).Error; err!=nil{
			return serializer.ParamErr("添加会议失败1",err)
		}
		return serializer.ParamGood("添加会议成功, meeting id = " + newMid)
	} else {
		return serializer.ParamErr("添加会议失败 : 已存在记录",nil)
	}
}
