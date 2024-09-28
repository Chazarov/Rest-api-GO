package main

import (
	"log"

	project "github.com/Chazarov/rest-app"
	"github.com/Chazarov/rest-app/pkg/handler"
	"github.com/Chazarov/rest-app/pkg/repository"
	"github.com/Chazarov/rest-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(project.Server)
	if err := srv.Run("8880", handlers.InitRoutes()); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

}
