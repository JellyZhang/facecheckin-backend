package api

import (
	"facecheckin/service"
	"github.com/gin-gonic/gin"
)

func CheckAdd(c *gin.Context) {
	var service service.CheckAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddCheck(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func CheckStatistic(c *gin.Context) {
	var service service.CheckstatisticService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetStatic(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func CheckFace(c *gin.Context) {
	var service service.BaiduFaceService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
