package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vavas/workchan/app/config"
	"github.com/vavas/workchan/app/server"
	"github.com/vavas/workchan/pkg/env"
)

func main() {

	appEnv := env.GetAppEnv()

	logger := logrus.StandardLogger()
	logger.Println("Environment: ", appEnv)
	err := env.LoadEnvFileIfNeeded(appEnv)
	if err != nil {
		logger.Fatalf("dotenv error: %v\n", err)
	}

	conf, err := config.Load(appEnv)

	router, err := server.ConfigureRouter(&server.RouterDeps{})
	if err != nil {
		logger.Fatalf("Could not initialize the Router: %v\n", err)
	}

	logger.Fatal(server.Start(router, conf.Server))
}
