package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uyupun/regret/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v0")
	api.GET("/service", handler.GetServices)
	api.POST("/service", handler.AddService)
	api.PATCH("/service", handler.EditService)
	api.DELETE("/service", handler.DeleteService)

	return e
}
