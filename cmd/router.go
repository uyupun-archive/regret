package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uyupun/regret/handler/admin"
	"github.com/uyupun/regret/handler/general"
	myMiddleware "github.com/uyupun/regret/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Re:gret")
	})

	apiVersion := "/api/v0"
	adminApi := e.Group(apiVersion+"/admin", myMiddleware.AppKeyVerification)
	generalApi := e.Group(apiVersion, myMiddleware.AccessTokenVerification)

	registerAdminRoutes(*adminApi)
	registerGeneralRoutes(*generalApi)

	return e
}

func registerAdminRoutes(adm echo.Group) {
	adm.GET("/service", admin.GetServices)
	adm.POST("/service", admin.AddService)
	adm.PATCH("/service", admin.EditService)
	adm.DELETE("/service", admin.DeleteService)

	adm.GET("/inquiry-validation", admin.GetInquiryValidation)

	adm.GET("/category", admin.GetCategories)
	adm.POST("/category", admin.AddCategory)
	adm.PATCH("/category", admin.EditCategory)
	adm.DELETE("/category", admin.DeleteCategory)
}

func registerGeneralRoutes(gen echo.Group) {
	gen.GET("/category", general.GetCategories)
	gen.POST("/inquiry", general.PostInquiry)
}
