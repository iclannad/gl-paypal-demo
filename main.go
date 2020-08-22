package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gl-paypal-demo/apires"
	"gl-paypal-demo/py"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	py.Init()
	e.GET("/pay/success", apires.Success)
	e.GET("/pay/cancel", apires.Cancel)
	e.GET("/pay/create", apires.Create)
	e.Start(":8778")
}

