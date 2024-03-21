package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wonderf00l/filmLib/internal/app"
)

//	@title			FilmLib API
//	@version		1.0
//	@description	API for films and actors library

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	serviceLogger, cfgFiles, err := app.Init()
	if err != nil {
		log.Fatal(err)
	}

	serviceCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	defer func() {
		if err = serviceLogger.Sync(); err != nil {
			log.Println("Sync service logger: ", err)
		}
	}()

	if err = app.Run(serviceCtx, serviceLogger, cfgFiles); err != nil {
		serviceLogger.Fatal(err)
	}
}
