package main

import (
	"fmt"
	"github.com/aibotsoft/forted-service/services/client"
	"github.com/aibotsoft/forted-service/services/handler"
	"github.com/aibotsoft/forted-service/services/server"
	"github.com/aibotsoft/forted-service/services/store"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/mig"
	"github.com/aibotsoft/micro/sqlserver"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	log := logger.New()
	log.Infow("Begin service", "config", cfg)
	db := sqlserver.MustConnect(cfg)
	err := mig.MigrateUp(cfg, log, db)
	if err != nil {
		log.Fatal(err)
	}
	cli := client.NewFortedClient(cfg, log)
	sto := store.NewStore(cfg, log, db)
	han := handler.NewHandler(cfg, log, cli, sto)

	s := server.NewServer(cfg, log, han)

	// Инициализируем GracefulStop
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errc <- s.Serve()
	}()
	defer func() {
		log.Debug("begin closing services")
		s.GracefulStop()
		_ = db.Close()
	}()

	log.Info("exit: ", <-errc)
}
