package main

import (
	"log"

	project "github.com/Chazarov/rest-app"
	"github.com/Chazarov/rest-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(project.Server)
	if err := srv.Run("8880", handlers.InitRoutes()); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

}
