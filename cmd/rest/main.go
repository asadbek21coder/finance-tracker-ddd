package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/asadbek21coder/fintracker2"
	"github.com/asadbek21coder/fintracker2/config"
	"github.com/asadbek21coder/fintracker2/internal/db"
	"github.com/asadbek21coder/fintracker2/internal/delivery/rest/router"

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

	db, err := db.ConnectToDb(cfg)

	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}

	// repos := adapters.NewUserRepository(db)
	// services := app.NewUserUseCase(db, repos)
	// router.InitRoutes(db)

	srv := new(fintracker2.Server)
	go func() {
		if err := srv.Run(os.Getenv("RPC_PORT"), router.InitRoutes(db)); err != nil {
			log.Fatal("error occured while running http server: ", err)
		}
	}()

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
