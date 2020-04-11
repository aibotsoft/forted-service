package main

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T) {
	got := Simple()
	cfg := config.New()
	log := logger.New()
	log.Infow("Begin service", "config", cfg)
	assert.Equal(t, 1, got)
}

func Test_formSourceURL(t *testing.T) {
	cfg := config.New()
	formSourceURL(cfg)
}

func Test_migrateUp(t *testing.T) {
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnect(cfg)
	migrateUp(cfg, log, db)

}
