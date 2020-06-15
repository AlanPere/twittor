package main

import (
	"log"

	"github.com/AlanPere/twittor/bd"
	"github.com/AlanPere/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}