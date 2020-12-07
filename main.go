package main

import (
	"log"

	"github.com/twittor/db"
	"github.com/twittor/handler"
)

func main() {

	if db.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la BD")
		return
	}

	handler.Manejadores()

}
