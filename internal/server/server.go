package server

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetServer(cfg *config.Config, a *models.App) *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())

	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.POST("/upload", a.UploadPhotoHandler)
	apiV1.GET("/photo", a.GetPhotoListHandler)
	apiV1.GET("/photo/:id", a.GetPhotoHandler)

	apiV1.GET("/show/preview/:id", a.ShowPreviewHandler)
	apiV1.GET("/show/image/:id", a.ShowPhotoHandler)
	apiV1.DELETE("/photo/:id", a.DelPhotoHandler)

	return e
}
