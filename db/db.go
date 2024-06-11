package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Estudante struct {
	gorm.Model
	Nome  string
	CPF   int
	Email string
	Idade int
	Ativo bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Estudante{})

	return db
}

func AddEstudante() {
	db := Init()

	estudante := Estudante{
		Nome:  "Dornas",
		CPF:   12345,
		Email: "dornas@gmail.com",
		Idade: 28,
		Ativo: true,
	}

	result := db.Create(&estudante)
	if result.Error != nil {
		fmt.Println("Erro para criar estudante")
	}

	fmt.Println("Estudante criado")
}
