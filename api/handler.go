package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dornascarol/API-estudantes/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (api *API) getEstudantes(c echo.Context) error {
	estudantes, err := api.DB.GetEstudantes()
	if err != nil {
		return c.String(http.StatusNotFound, "Falha em obter os estudantes")
	}

	listaEstudantes := map[string][]schemas.RespostaEstudante{"Estudantes": schemas.NovaResposta(estudantes)}

	return c.JSON(http.StatusOK, listaEstudantes)
}

func (api *API) createEstudante(c echo.Context) error {
	estudanteSolici := EstudanteSolicita{}
	if err := c.Bind(&estudanteSolici); err != nil {
		return err
	}

	if err := estudanteSolici.Validação(); err != nil {
		log.Error().Err(err).Msgf("[api] erro ao validar estrutura")
		return c.String(http.StatusBadRequest, "Erro para validar estudante")
	}

	estudante := schemas.Estudante{
		Nome:  estudanteSolici.Nome,
		CPF:   estudanteSolici.CPF,
		Email: estudanteSolici.Email,
		Idade: estudanteSolici.Idade,
		Ativo: *estudanteSolici.Ativo,
	}

	if err := api.DB.AddEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Erro para cadastrar estudante")
	}

	return c.JSON(http.StatusOK, estudante)
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

	recebeEstudante := schemas.Estudante{}
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter estudante")
	}

	estudante, err := api.DB.GetEstudante(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Estudante não encontrado")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao obter estudante")
	}

	if err := api.DB.DeleteEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Falha ao deletar estudante")
	}

	return c.JSON(http.StatusOK, estudante)
}

func atualizaInfoEstudante(recebeEstudante, estudante schemas.Estudante) schemas.Estudante {
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
