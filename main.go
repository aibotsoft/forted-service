package main

import (
	"github.com/aibotsoft/micro/logger"
)

func main() {
	log := logger.New()
	log.Debug("hello world")
}
