package main

import (
	"log"

	project "github.com/Chazarov/rest-app"
	"github.com/Chazarov/rest-app/pkg/handler"
	"github.com/Chazarov/rest-app/pkg/repository"
	"github.com/Chazarov/rest-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(project.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { /// Check GetString("port")!!!
		log.Fatalf("error initializing configs: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
