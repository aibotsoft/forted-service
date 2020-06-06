package main

import (
	"fmt"
	"github.com/aibotsoft/forted-service/services/client"
	"github.com/aibotsoft/forted-service/services/handler"
	"github.com/aibotsoft/forted-service/services/server"
	"github.com/aibotsoft/forted-service/services/store"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
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
	log.Infow("Begin service", "config", cfg.Service)
	conf := config_client.New(cfg, log)

	db := sqlserver.MustConnectX(cfg)
	err := mig.MigrateUp(cfg, log, db)
	if err != nil {
		log.Fatal(err)
	}
	cli := client.NewFortedClient(cfg, log, conf)

	sto := store.NewStore(cfg, log, db)
	han := handler.New(cfg, log, cli, sto, conf)

	s := server.NewServer(cfg, log, han)

	// Инициализируем Close
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() { errc <- s.Serve() }()
	defer func() { s.Close() }()
	log.Info("exit: ", <-errc)
}
