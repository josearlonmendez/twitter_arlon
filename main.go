package main

import (
	"log"

	"github.com/josearlonmendez/twitter_arlon/bd"
	"github.com/josearlonmendez/twitter_arlon/handlers"
)

func main() {
	if bd.ChequeoConnections() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
