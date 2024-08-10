package api

import (
	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"
	rabbitmq "api_gateway/pkg/messege_brokers/rabbitMQ"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title LinguaLeap
// @version 1.0
// @description Something big

// @contact.url http://www.support_me_with_smile

// @BasePath /
func NewRouter(log logger.ILogger, services client.IServiceManager, rabbitMQProducer rabbitmq.IRabbitMQProducer, iKafka kafka.IKafka) *gin.Engine {
	handlerV1 := v1.NewHandlerV1(services, log, rabbitMQProducer, iKafka)

	r := gin.Default()

	r.GET("swagger/*any", swagger.WrapHandler(swaggerFile.Handler))

	users := r.Group("/users")
	{
		users.GET("/profile", handlerV1.GetUserProfile)
		users.PUT("/update", handlerV1.UpdateUserProfile)
		users.PUT("/password", handlerV1.ChangePassword)
	}

	return r
}
