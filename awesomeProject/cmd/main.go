package main

import (
	"awesomeProject/config"
	"awesomeProject/logger"
	"awesomeProject/service"
	"awesomeProject/storage"
	"awesomeProject/transport"
	"awesomeProject/transport/handlers"
	"awesomeProject/transport/middleware"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

//	@title			OneLab HomeWork
//	@version		1.0
//	@description	API service for User, Book Storage.
//	@description	Where they can create, retrieve, update, delete books.
//  @description	And can rent these books
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	conf := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	l, err := logger.Init(conf)
	if err != nil {
		return fmt.Errorf("cannot init logger: %w", err)
	}
	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)

	storage, err := storage.NewStorage(l, ctx, conf)
	if err != nil {
		return err
	}

	defer cancel()

	gracefulShutdown(cancel)

	service, err := service.NewService(l, storage)
	if err != nil {
		return err
	}

	mid := middleware.NewJWTAuth(conf)

	handler := handlers.NewHandlers(l, conf, storage, service, mid)

	srv := transport.NewServer(handler, conf, mid)

	l.Info("Start server")
	err = srv.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}

func gracefulShutdown(ctx context.CancelFunc) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		log.Println(<-done)
		fmt.Println("Gracefully shutdown")
		ctx()
	}()
}
