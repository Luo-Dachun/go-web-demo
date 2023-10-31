package system

import (
	"github.com/gin-gonic/gin"
	v1 "web-demo/api/v1"
)

type BaseRouter struct{}

func (r *BaseRouter) InitBaseRouter(router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
	}
	return baseRouter
}
