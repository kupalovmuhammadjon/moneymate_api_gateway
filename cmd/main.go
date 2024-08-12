package main

import (
	"api_gateway/api"
	"api_gateway/configs"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"

	"github.com/casbin/casbin/v2"

	"go.uber.org/zap"
)

func main() {
	config := configs.Load()

	logger := logger.NewLogger(config.ServiceName, config.LoggerLevel, config.LogPath)

	services, err := client.NewGrpcClients(config)
	if err != nil {
		logger.Fatal("Failed make client connections ", zap.Error(err))
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer("./configs/model.conf", "./configs/policy.csv")
	if err != nil {
		logger.Error("Error while loading model and policy", zap.Error(err))
		return
	}

	iKafka, err := kafka.NewIKafka()
	if err != nil {
		logger.Fatal("Failed to create Kafka producer and consumer", zap.Error(err))
		return
	}
	defer iKafka.Close()

	router := api.NewRouter(logger, services, iKafka, casbinEnforcer)

	logger.Info("Fiber router is running..")
	err = router.Listen(config.ApiGatewayHttpHost + config.ApiGatewayHttpPort)
	if err != nil {
		logger.Fatal("Fiber router failed to run", zap.Error(err))
		return
	}
}
