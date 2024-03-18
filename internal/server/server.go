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

	routinesMaster := e.Group("/routinesMaster")
	api := routinesMaster.Group("/api")
	apiV1 := api.Group("/v1")

	// /routinesMaster/api/v1/ ....
	apiV1.GET("/routine/:id", a.GetRoutineHandler)
	apiV1.POST("/add/:id", a.AddCase)
	apiV1.DELETE("/delete/:id/:case_id", a.DeleteCase)
	apiV1.GET("/sound/:path", a.GetSoundHandler)

	return e
}
