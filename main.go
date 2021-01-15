package main

import (
	"fmt"
	"github.com/aibotsoft/forted-service/services/handler"
	"github.com/aibotsoft/forted-service/services/server"
	"github.com/aibotsoft/forted-service/services/store"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/config_client"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/subosito/gotenv"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cfg := config.New()
	log := logger.New()

	log.Infow("Begin service", "config", cfg.Service)
	conf := config_client.New(cfg, log)

	db := sqlserver.MustConnectX(cfg)
	//err := mig.MigrateUp(cfg, log, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//cli := client.NewFortedClient(cfg, log, conf)

	sto := store.NewStore(cfg, log, db)
	h := handler.New(cfg, log, sto, conf)
	h.Publish("ping", "ping")
	s := server.NewServer(cfg, log, h)

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
