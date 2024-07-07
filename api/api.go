package api

import (
	"github.com/dornascarol/API-estudantes/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/dornascarol/sistema-estudantes/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.EstudanteManipula
}

// @title API Estudantes
// @version 1.0
// @description This is a sample server API Estudantes
// @host localhost:8080
// @BasePath /
// @schemes http

func NovoServidor() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	estudanteDB := db.NewEstudanteManipula(database)

	return &API{
		Echo: e,
		DB:   estudanteDB,
	}
}

func (api *API) ConfigRotas() {
	api.Echo.GET("/estudantes", api.getEstudantes)
	api.Echo.POST("/estudantes", api.createEstudante)
	api.Echo.GET("/estudantes/:id", api.getEstudante)
	api.Echo.PUT("/estudantes/:id", api.updateEstudante)
	api.Echo.DELETE("/estudantes/:id", api.deleteEstudante)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (api *API) ComecaServidor() error {
	return api.Echo.Start(":8080")
}
