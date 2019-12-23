package server

import (
	"facecheckin/api"
	"facecheckin/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/reg", api.UserRegister)

		//用户登录
		v1.POST("user/login", api.UserLogin)

		// user update
		v1.POST("user/update", api.UserUpdate)

		// user info
		v1.GET("user/info", api.UserInfo)

		//meeting detail
		v1.GET("meeting/detail", api.MeetingDetial)

		//meeting add
		v1.POST("meeting/add", api.MeetingAdd)

		//meeting update
		v1.POST("meeting/update", api.MeetingUpdate)

		//meeting join
		v1.POST("meeting/join", api.MeetingJoin)

		//meeting leave
		v1.POST("meeting/leave", api.MeetingLeave)

		// chech add
		v1.POST("check/add", api.CheckAdd)

		// check statistic
		v1.POST("check/statistic", api.CheckStatistic)
		// 需要登录保护的
		//auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		//{
		//	// User Routing
		//	auth.GET("user/me", api.UserMe)
		//	auth.DELETE("user/logout", api.UserLogout)
		//}
	}
	return r
}
