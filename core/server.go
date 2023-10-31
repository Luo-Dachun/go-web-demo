package core

import (
	"web-demo/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 为了之后做其他操作
	Router := initialize.Routers()

	if err := Router.Run(); err != nil {
		panic(err.Error())
	}

	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	//s := initServer(address, Router)
	//global.GVA_LOG.Error(s.ListenAndServe().Error())
}
