package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dornascarol/API-estudantes/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (api *API) getEstudantes(c echo.Context) error {
	estudantes, err := api.DB.GetEstudantes()
	if err != nil {
		return c.String(http.StatusNotFound, "Falha em obter os estudantes")
	}

	return c.JSON(http.StatusOK, estudantes)
}

func (api *API) createEstudante(c echo.Context) error {
	estudante := db.Estudante{}
	if err := c.Bind(&estudante); err != nil {
		return err
	}

	if err := api.DB.AddEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Erro para cadastrar estudante")
	}

	return c.String(http.StatusOK, "Estudante cadastrado")
}

func (api *API) getEstudante(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter ID do estudante")
	}

	estudante, err := api.DB.GetEstudante(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Estudante não encontrado")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter estudante")
	}

	return c.JSON(http.StatusOK, estudante)
}

func (api *API) updateEstudante(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter estudante")
	}

	recebeEstudante := db.Estudante{}
	if err := c.Bind(&recebeEstudante); err != nil {
		return err
	}

	atualizandoEstudante, err := api.DB.GetEstudante(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Estudante não encontrado")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter estudante")
	}

	estudante := atualizaInfoEstudante(recebeEstudante, atualizandoEstudante)

	if err := api.DB.UpdateEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao salvar estudante atualizado")
	}

	return c.JSON(http.StatusOK, estudante)
}

func (api *API) deleteEstudante(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Deleta %s estudante", id)
	return c.String(http.StatusOK, deleteStud)
}

func atualizaInfoEstudante(recebeEstudante, estudante db.Estudante) db.Estudante {
	if recebeEstudante.Nome != "" {
		estudante.Nome = recebeEstudante.Nome
	}
	if recebeEstudante.CPF > 0 {
		estudante.CPF = recebeEstudante.CPF
	}
	if recebeEstudante.Email != "" {
		estudante.Email = recebeEstudante.Email
	}
	if recebeEstudante.Idade > 0 {
		estudante.Idade = recebeEstudante.Idade
	}
	if recebeEstudante.Ativo != estudante.Ativo {
		estudante.Ativo = recebeEstudante.Ativo
	}
	return estudante
}
