package main

import (
	"github.com/KSkun/doub-freshman/config"
	"github.com/KSkun/doub-freshman/controller"
	middleware2 "github.com/KSkun/doub-freshman/middleware"
	"github.com/KSkun/doub-freshman/router"
	"github.com/KSkun/doub-freshman/util"
	. "github.com/KSkun/doub-freshman/util/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ok, err := util.ParseFlag()
	if err != nil {
		Logger.Fatal(err)
	}

	if !ok {
		return
	}

	e := echo.New()

	// 自定义未处理错误的handler
	e.HTTPErrorHandler = controller.HTTPErrorHandler

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = middleware2.GetValidator()
	err = middleware2.InitBeforeStart(e)
	if err != nil {
		Logger.Fatal(err)
	}

	gAPI := e.Group(config.C.App.Prefix)
	router.InitRouter(gAPI)

	Logger.Fatal(e.Start(config.C.App.Addr))
}
