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