package main

import (
	handler "avito-third/internal/handlers"
	"avito-third/internal/repository"
	"avito-third/internal/service"
	"avito-third/pkg/client/postgresql"
	"avito-third/pkg/server"
	"context"
	_ "context"
	"os"
	"os/signal"
	_ "os/signal"
	"syscall"
	_ "syscall"

	_ "github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title AvitoTech
// @version 1.0
// @description API Server for segment distribution

// @host localhost:8000
// @BasePath /

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := postgresql.NewClient(postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Database: viper.GetString("db.database"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()

	go func() {
		if err := service.Start(*repos); err != nil {
			logrus.Fatalf("checker not launched..", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
