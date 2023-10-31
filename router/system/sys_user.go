package system

import (
	"github.com/gin-gonic/gin"
	v1 "web-demo/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userRouterWithoutRecord := router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register) // 注册账号
		//userRouter.POST("login", baseApi.Login)       // 登录账号
	}

	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)
	}
}
