package main

import (
	"log"

	"github.com/dornascarol/API-estudantes/api"
)

func main() {
	servidor := api.NovoServidor()

	servidor.ConfigRotas()

	if err := servidor.ComecaServidor(); err != nil {
		log.Fatal(err)
	}
}
