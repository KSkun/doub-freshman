package router

import (
	"github.com/KSkun/doub-freshman/controller"
	"github.com/labstack/echo/v4"
)

// InitRouter 初始化所有路由，可以每个路由分函数分文件写，方便之后维护
func InitRouter(g *echo.Group) {
	initIndexRouter(g)
}

func initIndexRouter(g *echo.Group) {
	g.POST("/new", controller.HandlerNewGame)
	g.PUT("/option", controller.HandlerSelectOption)
	g.PUT("/stage", controller.HandlerSelectStage)
	g.POST("/stage", controller.HandlePostStage)
}
