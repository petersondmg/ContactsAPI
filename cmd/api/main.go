package main

import (
	"capi/cmd/api/internal/router"
	"capi/domain/repository"
	"capi/domain/service"
	"capi/pkg/config"
	"capi/pkg/database"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("can't load config", zap.Error(err))
	}

	srv, err := initService(cfg)
	if err != nil {
		log.Fatal("can't init service", zap.Error(err))
	}

	httpServer := http.Server{
		Addr:         cfg.API.Addr,
		Handler:      router.New(srv, log, cfg.API.JWTSecret),
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 30,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	go func() {
		// listen to interrupt signal and shutdown gracefully
		<-done

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error("error on server shutdown", zap.Error(err))
		}
	}()

	log.Info("starting api server", zap.String("addr", cfg.API.Addr))
	if err := httpServer.ListenAndServe(); err != nil {
		log.Error("error starting server", zap.Error(err))
	}
}

func initService(c *config.Config) (*service.Service, error) {
	dbVarejao, err := database.ConnectPG(c.VarejaoDB.User, c.VarejaoDB.Pass, c.VarejaoDB.Database, c.VarejaoDB.Addr)
	if err != nil {
		return nil, fmt.Errorf("error opening varejao pgsql: %w", err)
	}

	dbMacapa, err := database.ConnectMySQL(c.MacapaDB.User, c.MacapaDB.Pass, c.MacapaDB.Database, c.MacapaDB.Addr)
	if err != nil {
		return nil, fmt.Errorf("error opening macapa mysql: %w", err)
	}

	macapaContactRepo := repository.NewMacapaContact(dbMacapa)
	varejaoContactRepo := repository.NewVarejaoContact(dbVarejao)

	srv := service.New(macapaContactRepo, varejaoContactRepo)

	return srv, nil
}
