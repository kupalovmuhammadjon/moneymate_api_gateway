package v1

import (
	"api_gateway/api/handlers/models"
	checker  "api_gateway/pkg/jwt"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
)

type HandlerV1 struct {
	services client.IServiceManager
	log      logger.ILogger
	iKafka   kafka.IKafka
}

func NewHandlerV1(services client.IServiceManager, logger logger.ILogger, iKafka kafka.IKafka) *HandlerV1 {
	return &HandlerV1{
		services: services,
		log:      logger,
		iKafka:   iKafka,
	}
}

func handleResponse(ctx *fiber.Ctx, log logger.ILogger, msg string, statusCode int, data interface{}) error {
	var resp models.Response

	switch {
	case statusCode >= 200 && statusCode < 300:
		resp.Description = "OK"
		log.Info("Response OK", logger.String("msg", msg), logger.Int("status", statusCode))
	case statusCode == 400:
		resp.Description = "Bad Request"
		log.Warn("Bad Request", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	case statusCode == 401:
		resp.Description = "Unauthorized"
		log.Warn("Unauthorized", logger.String("msg", msg), logger.Int("status", statusCode))
	case statusCode >= 500:
		resp.Description = "Internal Server Error"
		log.Error("Internal Server Error", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	default:
		resp.Description = "Unknown Error"
		log.Error("Unknown Error", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	return ctx.Status(statusCode).JSON(resp)
}

func getUserInfoFromToken(ctx *fiber.Ctx) (*models.UserInfoFromToken, error) {

	var (
		token  string
		claims jwt.MapClaims
		resp   = models.UserInfoFromToken{}
		err    error
	)

	token = ctx.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("authorization is requeired")
	}

	claims, err = checker.ExtractClaims(token)
	if err != nil {
		return nil, err
	}

	resp.Id = claims["user_id"].(string)
	resp.Username = claims["username"].(string)
	resp.Email = claims["email"].(string)
	resp.FirstName = claims["first_name"].(string)
	resp.LastName = claims["last_name"].(string)
	resp.Role = claims["role"].(string)

	return &resp, nil
}
