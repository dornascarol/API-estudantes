package api

import (
	"fmt"
	"net/http"

	"github.com/dornascarol/API-estudantes/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NovoServidor() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) ConfigRotas() {
	api.Echo.GET("/estudantes", getEstudantes)
	api.Echo.POST("/estudantes", createEstudante)
	api.Echo.GET("/estudantes/:id", getEstudante)
	api.Echo.PUT("/estudantes/:id", updateEstudante)
	api.Echo.DELETE("/estudantes/:id", deleteEstudante)
}

func (api *API) ComecaServidor() error {
	return api.Echo.Start(":8080")
}

func getEstudantes(c echo.Context) error {
	estudantes, err := db.GetEstudantes()
	if err != nil {
		return c.String(http.StatusNotFound, "Falha em obter os estudantes")
	}

	return c.JSON(http.StatusOK, estudantes)
}

func createEstudante(c echo.Context) error {
	estudante := db.Estudante{}
	if err := c.Bind(&estudante); err != nil {
		return err
	}

	if err := db.AddEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Erro para cadastrar estudante")
	}

	return c.String(http.StatusOK, "Estudante cadastrado")
}

func getEstudante(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Pega %s estudante", id)
	return c.String(http.StatusOK, getStud)
}

func updateEstudante(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Atualiza %s estudante", id)
	return c.String(http.StatusOK, updateStud)
}

func deleteEstudante(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Deleta %s estudante", id)
	return c.String(http.StatusOK, deleteStud)
}
