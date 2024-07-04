package api

import "fmt"

type EstudanteSolicita struct {
	Nome  string `json:"nome"`
	CPF   int    `json:"cpf"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
	Ativo *bool  `json:"ativo"`
}

func erroParamSolicita(param, typ string) error {
	return fmt.Errorf("parâmetro '%s' do tipo '%s' é obrigatório", param, typ)
}

func (s *EstudanteSolicita) Validação() error {
	if s.Nome == "" {
		return erroParamSolicita("nome", "string")
	}
	if s.CPF == 0 {
		return erroParamSolicita("cpf", "int")
	}
	if s.Email == "" {
		return erroParamSolicita("email", "string")
	}
	if s.Idade == 0 {
		return erroParamSolicita("idade", "int")
	}
	if s.Ativo == nil {
		return erroParamSolicita("ativo", "bool")
	}

	return nil
}
