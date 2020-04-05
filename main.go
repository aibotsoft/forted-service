package main

import (
	"fmt"
	"github.com/aibotsoft/forted-service/services/server"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
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
	s := server.NewServer(cfg, log, db)

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

func Simple() int {
	log := logger.New()
	log.Debug("hello world")
	log.Debug("hello world another time")
	return 1
}
