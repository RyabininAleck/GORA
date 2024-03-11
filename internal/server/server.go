package server

import (
	"GORAbackend/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func GetServer(cfg *config.Config) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	// Routes
	apiV1.POST("/upload", loadPhotoHandler)
	apiV1.GET("/photo", getPhotoListHandler)
	apiV1.GET("/photo/:id", getPhotoHandler)
	apiV1.DELETE("/photo/:id", delPhotoHandler)

	return e

}

func loadPhotoHandler(c echo.Context) error {
	// todo предусмотреть 413 Payload Too Large

	_, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "loadPhotoHandler")
}

func getPhotoListHandler(c echo.Context) error {
	return c.String(http.StatusOK, "getPhotoListHandler")
}

func getPhotoHandler(c echo.Context) error {
	id := c.Param("id")

	// Делаем что-то с полученным идентификатором, например, возвращаем его в ответе
	return c.String(http.StatusOK, "getPhotoHandler: Photo ID: "+id)

}

func delPhotoHandler(c echo.Context) error {
	id := c.Param("id")

	// Делаем что-то с полученным идентификатором, например, возвращаем его в ответе
	return c.String(http.StatusOK, "delPhotoHandler: Photo ID: "+id)

}
