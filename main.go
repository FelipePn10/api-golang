package main

import (
	"github.com/FelipePn10/api-golang/config"
	"github.com/FelipePn10/api-golang/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error:: %v", err)
		return
	}

	router.Initialize()
}
