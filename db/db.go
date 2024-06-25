package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type EstudanteManipula struct {
	DB *gorm.DB
}

type Estudante struct {
	gorm.Model
	Nome  string `json:"nome"`
	CPF   int    `json:"cpf"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
	Ativo bool   `json:"ativo"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Estudante{})

	return db
}

func NewEstudanteManipula(db *gorm.DB) *EstudanteManipula {
	return &EstudanteManipula{DB: db}
}

func (s *EstudanteManipula) AddEstudante(estudante Estudante) error {

	if result := s.DB.Create(&estudante); result.Error != nil {
		return result.Error
	}

	fmt.Println("Estudante criado")
	return nil
}

func (s *EstudanteManipula) GetEstudantes() ([]Estudante, error) {
	estudantes := []Estudante{}

	err := s.DB.Find(&estudantes).Error
	return estudantes, err
}
