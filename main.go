package main

import (
	"fmt"
	"web-demo/core"
	"web-demo/global"
	"web-demo/initialize"
)

func main() {
	// 初始化Viper
	global.GVA_VP = core.Viper()
	// gorm连接数据库
	global.GVA_DB = initialize.Gorm()
	core.RunWindowsServer()
	fmt.Println("Hello World")
}
