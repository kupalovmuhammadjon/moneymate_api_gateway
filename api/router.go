// @title MoneyMate
// @version 1.0
// @description Something big
// @BasePath /
package api

import (
	_ "api_gateway/api/docs"
	"api_gateway/api/handlers/middleware"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/pkg/messege_brokers/kafka"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/segmentio/encoding/json"
	swagger "github.com/swaggo/fiber-swagger"
)

// @title MoneyMate
// @version 1.0
// @description Something big
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func NewRouter(log logger.ILogger, services client.IServiceManager, iKafka kafka.IKafka, casbinEnforcer *casbin.Enforcer) *fiber.App {
	handlerV1 := v1.NewHandlerV1(services, log, iKafka)

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

	users := router.Group("/users", middleware.JWTMiddleware(casbinEnforcer))
	{
		users.Get("/profile", handlerV1.GetUserProfile)
		users.Put("/update", handlerV1.UpdateUserProfile)
		users.Put("/password", handlerV1.ChangePassword)
	}

	accounts := router.Group("/accounts", middleware.JWTMiddleware(casbinEnforcer))
	{
		accounts.Post("/create", handlerV1.CreateAccount)
		accounts.Get("/all", handlerV1.GetAllAccounts)
		accounts.Get("/:id", handlerV1.GetAccountById)
		accounts.Put("/:id/update", handlerV1.UpdateAccount)
		accounts.Delete("/:id/delete", handlerV1.DeleteAccount)
	}

	budgets := router.Group("/budgets", middleware.JWTMiddleware(casbinEnforcer))
	{
		budgets.Post("/create", handlerV1.CreateBudget)
		budgets.Get("/all", handlerV1.GetAllBudgets)
		budgets.Get("/:id", handlerV1.GetBudgetById)
		budgets.Put("/:id/update", handlerV1.UpdateBudget)
		budgets.Delete("/:id/delete", handlerV1.DeleteBudget)
	}

	categories := router.Group("/categories", middleware.JWTMiddleware(casbinEnforcer))
	{
		categories.Post("/create", handlerV1.CreateCategory)
		categories.Get("/all", handlerV1.GetAllCategories)
		categories.Get("/:id", handlerV1.GetCategoryById)
		categories.Put("/:id/update", handlerV1.UpdateCategory)
		categories.Delete("/:id/delete", handlerV1.DeleteCategory)
	}

	goals := router.Group("/goals", middleware.JWTMiddleware(casbinEnforcer))
	{
		goals.Post("/create", handlerV1.CreateGoal)
		goals.Get("/:id", handlerV1.GetGoalById)
		goals.Get("/all", handlerV1.GetAllGoals)
		goals.Put("/:id/update", handlerV1.UpdateGoal)
		goals.Delete("/:id/delete", handlerV1.DeleteGoal)
	}

	transactions := router.Group("/transactions", middleware.JWTMiddleware(casbinEnforcer))
	{
		transactions.Post("/create", handlerV1.CreateTransaction)
		transactions.Get("/:id", handlerV1.GetTransactionById)
		transactions.Get("/all", handlerV1.GetAllTransactions)
		transactions.Put("/:id/update", handlerV1.UpdateTransaction)
		transactions.Delete("/:id/delete", handlerV1.DeleteTransaction)
	}

	return router
}
