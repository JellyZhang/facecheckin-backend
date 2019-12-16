package api

import (
	"facecheckin/service"
	"github.com/gin-gonic/gin"
)

func CheckAdd(c* gin.Context) {
	var service service.CheckAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddCheck(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
