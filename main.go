package main

import (
	"log"

	"github.com/grak0s/twittor/db"
	"github.com/grak0s/twittor/handler"
)

func main() {

	if db.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la BD")
		return
	}

	handler.Manejadores()

}
