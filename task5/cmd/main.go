package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task5"
	"task5/pkg/handler"
	"task5/pkg/repository"
	"task5/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil{
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err!=nil{
		log.Fatalf("Failed initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(task5.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"),handlers.InitRoutes()); err != nil{
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	log.Print("App shutting down")

	if err := srv.ShutDown(context.Background()); err !=nil{
		log.Fatalf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil{
		log.Fatalf("Error occured on db connection close: %s", err.Error())
	}
}

func InitConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
