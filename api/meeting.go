package api

import (
	"facecheckin/service"
	"github.com/gin-gonic/gin"
)

func MeetingDetial(c *gin.Context) {
	var service service.MeetingDetialService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetDetail(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func MeetingAdd(c *gin.Context) {
	var service service.MeetingAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddMeeting(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func MeetingUpdate(c *gin.Context) {
	var service service.MeetingUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateMeeting(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func MeetingJoin(c *gin.Context) {
	var service service.MeetingJoinService
	if err := c.ShouldBind(&service); err == nil {
		res := service.JoinMeeting(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func MeetingLeave(c *gin.Context) {
	var service service.MeetingLeaveService
	if err := c.ShouldBind(&service); err == nil {
		res := service.LeaveMeeting(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func MeetingList(c *gin.Context) {
	var service service.MeetingListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
