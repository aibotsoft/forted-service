package main

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
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
