package main

import (
	"log"
	"twitton/db"
	"twitton/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
