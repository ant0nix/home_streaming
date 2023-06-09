package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	streaming "github.com/ant0nix/home_streaming"
	"github.com/ant0nix/home_streaming/internal/delivery"
	"github.com/ant0nix/home_streaming/internal/usecase"
	torrentClient "github.com/ant0nix/home_streaming/internal/usecase/torrent-client"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Error with config initializing! Error:%s", err.Error())
	}

	cfg := torrentClient.NewConfig()
	tclient, err := torrentClient.New(*cfg)
	if err != nil {
		log.Fatalf("Error with starting torrent client: %s", err.Error())
	}
	defer tclient.Client.Close()
	if err != nil {
		log.Println(err)
	}
	uc := usecase.New(tclient)
	handler := delivery.NewHandler(uc)
	srv := new(streaming.Server)
	go func() {
		if err := srv.Start(viper.GetString("port"), handler.InitRouters()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server Shutting Down")

	if err := srv.Stop(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
