package router

import (
	"web-demo/router/example"
	"web-demo/router/system"
)

type RouteGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouteGroup)
