package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/estudantes", getEstudantes)
	e.POST("/estudantes", createEstudante)
	e.GET("/estudantes/:id", getEstudante)
	e.PUT("/estudantes/:id", updateEstudante)
	e.DELETE("/estudantes/:id", deleteEstudante)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getEstudantes(c echo.Context) error {
	return c.String(http.StatusOK, "Lista de todos os estudantes")
}

func createEstudante(c echo.Context) error {
	return c.String(http.StatusOK, "Cadastro um estudante")
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
