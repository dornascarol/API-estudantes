package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
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

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Estudante{})

	return db
}

func AddEstudante(estudante Estudante) error {
	db := Init()

	result := db.Create(&estudante)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Estudante criado")
	return nil
}

func GetEstudantes() ([]Estudante, error) {
	estudantes := []Estudante{}

	db := Init()
	err := db.Find(&estudantes).Error
	return estudantes, err
}
