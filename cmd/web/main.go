package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"github.com/vavas/workchan/app/config"
	"github.com/vavas/workchan/app/server"
	"github.com/vavas/workchan/pkg/db"
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

	dbConns, err := db.Load(conf.Database)
	if err != nil {
		logger.Fatalf("Could not connect to the DB: %v\n", err)
	}

	defer dbConns.Close()

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer kafkaProducer.Close()

	router, err := server.ConfigureRouter(&server.RouterDeps{
		DbConns:       dbConns,
		KafkaProducer: kafkaProducer,
	})
	if err != nil {
		logger.Fatalf("Could not initialize the Router: %v\n", err)
	}

	logger.Fatal(server.Start(router, conf.Server))
}
