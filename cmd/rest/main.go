package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/asadbek21coder/fintracker2/config"
	"github.com/asadbek21coder/fintracker2/internal/adapters"
	"github.com/asadbek21coder/fintracker2/internal/app"
	"github.com/asadbek21coder/fintracker2/internal/db"
	"github.com/asadbek21coder/fintracker2/internal/delivery/rest"
	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	dotenv := cast.ToString(config.GetOrReturnDefault("DOT_ENV_PATH", ".env"))
	cfg := config.Load(dotenv)
	// log := logger.New(cfg.LogLevel)

	db, err := db.ConnectToDb(cfg)

	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}

	repos := adapters.NewUserRepository(db)
	services := app.NewUserUseCase(db, repos)
	handlers := rest.NewHandler(services)

	srv := new(domain.Server)
	go func() {
		if err := srv.Run(os.Getenv("RPC_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatal("error occured while running http server: ", err.Error())
		}
	}()

	fmt.Print("Demo app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Print("Demo app Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("error occured on server shutting down: ", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatal("error occured on db connection close: ", err.Error())
	}
}
