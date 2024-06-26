package main

import (
	"github.com/dornascarol/API-estudantes/api"
	"github.com/rs/zerolog/log"
)

func main() {
	servidor := api.NovoServidor()

	servidor.ConfigRotas()

	if err := servidor.ComecaServidor(); err != nil {
		log.Fatal().Err(err).Msgf("Falha ao iniciar o servidor")
	}
}
