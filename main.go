package main

import (
	"github.com/aibotsoft/micro/logger"
)

func main() {
	log := logger.New()
	log.Debug("hello world")
}

func Simple() int {
	log := logger.New()
	log.Debug("hello world")
	log.Debug("hello world another time")
	return 1
}
