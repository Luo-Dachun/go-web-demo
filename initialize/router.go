package initialize

import (
	"github.com/gin-gonic/gin"
	"web-demo/global"
	"web-demo/middleware"
	"web-demo/router"
)

/*
Routers
初始化总路由
*/
func Routers() *gin.Engine {
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}
	Router := gin.New()

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	PublicGroup := Router.Group("v1")

	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("v1")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup) // 注册用户路由

		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	}
	return Router
}
