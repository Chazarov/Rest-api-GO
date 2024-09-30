package main

import (
	project "github.com/Chazarov/rest-app"
	"github.com/Chazarov/rest-app/pkg/handler"
	"github.com/Chazarov/rest-app/pkg/repository"
	"github.com/Chazarov/rest-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// @title REST API Documentation
// @version 1.0

// @host localhost:8080
// @BasePath /api

// @in header
// @name Authorization
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("!!! error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("!!! error with loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("!!! failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(project.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { /// Check GetString("port")!!!
		logrus.Fatalf("!!! error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
