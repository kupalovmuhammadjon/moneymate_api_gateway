package api

import (
	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"
	rabbitmq "api_gateway/pkg/messege_brokers/rabbitMQ"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/segmentio/encoding/json"
	swagger "github.com/swaggo/fiber-swagger"
)

// @title MoneyMate
// @version 1.0
// @description Something big

// @BasePath /
func NewRouter(log logger.ILogger, services client.IServiceManager, rabbitMQProducer rabbitmq.IRabbitMQProducer, iKafka kafka.IKafka) *fiber.App {
	handlerV1 := v1.NewHandlerV1(services, log, rabbitMQProducer, iKafka)

	router := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	cors := cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE",
		AllowHeaders:  "Authorization",
		ExposeHeaders: "Authorization",
		MaxAge:        12 * int(time.Hour),
	})

	router.Use(cors)

	router.Get("/swagger/*", swagger.WrapHandler)

	users := router.Group("/users")
	{
		users.Get("/profile", handlerV1.GetUserProfile)
		users.Put("/update", handlerV1.UpdateUserProfile)
		users.Put("/password", handlerV1.ChangePassword)
	}

	return router
}
