package v1

import (
	"api_gateway/api/handlers/models"
	"api_gateway/api/handlers/tokens"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"
	rabbitmq "api_gateway/pkg/messege_brokers/rabbitMQ"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type HandlerV1 struct {
	services         client.IServiceManager
	log              logger.ILogger
	rabbitMQProducer rabbitmq.IRabbitMQProducer
	iKafka           kafka.IKafka
}

func NewHandlerV1(services client.IServiceManager, logger logger.ILogger, rabbitMQProducer rabbitmq.IRabbitMQProducer, iKafka kafka.IKafka) *HandlerV1 {
	return &HandlerV1{
		services:         services,
		log:              logger,
		rabbitMQProducer: rabbitMQProducer,
		iKafka:           iKafka,
	}
}

func handleResponse(ctx *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {

	var (
		resp = models.Response{}
	)

	switch code := statusCode; {
	case code < 250:
		resp.Description = "OK"
		log.Info("~~~~> OK", logger.String("msg", msg), logger.Any("status", statusCode))
	case code == 401:
		resp.Description = "Unauthorized"
		log.Info("????? Unauthorized", logger.String("msg", msg), logger.Any("status", statusCode))
	case code == 500:
		resp.Description = "Bad Request"
		log.Info("!!!!! Bad Request", logger.String("msg", msg), logger.Any("status", statusCode), logger.Any("Error", data))
	default:
		resp.Description = "Internal Server Error"
		log.Info("!!!!! Internal Server Error", logger.String("msg", msg), logger.Any("status", statusCode), logger.Any("Error", data))
	}

	resp.StatisCode = statusCode
	resp.Data = data

	ctx.JSON(statusCode, resp)
}

func getUserInfoFromToken(ctx *gin.Context) (*models.UserInfoFromToken, error) {

	var (
		token  string
		claims jwt.MapClaims
		resp   = models.UserInfoFromToken{}
		err    error
	)

	token = ctx.GetHeader("Authorization")
	if token == "" {
		return nil, fmt.Errorf("authorization is requeired")
	}
	fmt.Println(token)

	claims, err = tokens.ExtractClaims(token)
	if err != nil {
		return nil, err
	}
	fmt.Println(claims)

	resp.Id = claims["user_id"].(string)
	resp.Username = claims["username"].(string)
	resp.Email = claims["email"].(string)
	resp.FullName = claims["full_name"].(string)

	return &resp, nil
}
