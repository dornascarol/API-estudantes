package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Estudante struct {
	gorm.Model
	Nome  string `json:"nome"`
	CPF   int    `json:"cpf"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
	Ativo bool   `json:"ativo"`
}

type RespostaEstudante struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Nome      string    `json:"nome"`
	CPF       int       `json:"cpf"`
	Email     string    `json:"email"`
	Idade     int       `json:"idade"`
	Ativo     bool      `json:"ativo"`
}

func NovaResposta(estudantes []Estudante) []RespostaEstudante {
	respostaEstudantes := []RespostaEstudante{}

	for _, estudante := range estudantes {
		respostaEstudante := RespostaEstudante{
			ID:        int(estudante.ID),
			CreatedAt: estudante.CreatedAt,
			UpdatedAt: estudante.UpdatedAt,
			Nome:      estudante.Nome,
			CPF:       estudante.CPF,
			Email:     estudante.Email,
			Idade:     estudante.Idade,
			Ativo:     estudante.Ativo,
		}
		respostaEstudantes = append(respostaEstudantes, respostaEstudante)
	}
	return respostaEstudantes
}
