package api

import (
	"github.com/dornascarol/API-estudantes/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.EstudanteManipula
}

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
}

func (api *API) ComecaServidor() error {
	return api.Echo.Start(":8080")
}
