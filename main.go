package main

import (
	"fmt"
	"github.com/aibotsoft/forted-service/services/server"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/mig"
	"github.com/aibotsoft/micro/sqlserver"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	"os"
	"os/signal"
	"syscall"
)

//func formSourceURL(cfg *config.Config) {
//	result := fmt.Sprintf("github://%s:%s@%s/%s/migrations", cfg.Migrate.User, cfg.Migrate.Token, cfg.Migrate.User, cfg.Service.Name)
//	log.Print(result)
//}

//func migrateUp(cfg *config.Config, log *zap.SugaredLogger, db *sql.DB) {
//	start := time.Now()
//	log.Info("begin db migration..")
//	instance, err := stub.WithInstance(db, &stub.Config{SchemaName: "dbo", DatabaseName: cfg.Mssql.Database})
//	if err != nil {
//		log.Fatal(err)
//	}
//	version, dirty, err := instance.Version()
//	if err != nil {
//		log.Fatalf("instance.Version error: ", err)
//	}
//	log.Info("current db version: ", version, ", is dirty: ", dirty)
//	sourceURL := fmt.Sprintf("github://%s:%s@%s/%s/migrations", cfg.Migrate.User, cfg.Migrate.Token, cfg.Migrate.User, cfg.Service.Name)
//
//	m, err := migrate.NewWithDatabaseInstance(sourceURL, "sqlserver", instance)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = m.Up()
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Info("migration took: ", time.Since(start).Milliseconds())
//}

func main() {
	cfg := config.New()
	log := logger.New()
	log.Infow("Begin service", "config", cfg)
	db := sqlserver.MustConnect(cfg)
	err := mig.MigrateUp(cfg, log, db)
	log.Error(err)
	return
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
