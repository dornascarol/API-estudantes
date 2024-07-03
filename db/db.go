package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dornascarol/API-estudantes/schemas"
)

type EstudanteManipula struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Falha ao inicializar o SQLite: %s", err.Error())
	}

	db.AutoMigrate(&schemas.Estudante{})

	return db
}

func NewEstudanteManipula(db *gorm.DB) *EstudanteManipula {
	return &EstudanteManipula{DB: db}
}

func (s *EstudanteManipula) AddEstudante(estudante schemas.Estudante) error {
	if result := s.DB.Create(&estudante); result.Error != nil {
		log.Error().Msg("Falha para cadastrar estudante")
		return result.Error
	}

	log.Info().Msg("Estudante cadastrado")
	return nil
}

func (s *EstudanteManipula) GetEstudantes() ([]schemas.Estudante, error) {
	estudantes := []schemas.Estudante{}
	err := s.DB.Find(&estudantes).Error
	return estudantes, err
}

func (s *EstudanteManipula) GetEstudante(id int) (schemas.Estudante, error) {
	var estudante schemas.Estudante
	err := s.DB.First(&estudante, id)
	return estudante, err.Error
}

func (s *EstudanteManipula) UpdateEstudante(updateEstudante schemas.Estudante) error {
	return s.DB.Save(&updateEstudante).Error
}

func (s *EstudanteManipula) DeleteEstudante(estudante schemas.Estudante) error {
	return s.DB.Delete(&estudante).Error
}
