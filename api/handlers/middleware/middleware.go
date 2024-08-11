package middleware

import (
	"api_gateway/api/handlers/models"
	"api_gateway/pkg/jwt"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func JWTMiddleware(enforcer *casbin.Enforcer) func(ctx *fiber.Ctx) error {

	casbinPermission := casbinPermission{
		enforcer: enforcer,
	}

	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get("Authorization")
		if auth == "" {
			return ctx.Status(401).JSON(models.Response{
				StatusCode:  401,
				Description: "no Authorization header",
				Data:        nil,
			})
		}

		claims, err := jwt.ExtractClaims(auth)
		if err != nil {
			return ctx.Status(401).JSON(models.Response{
				StatusCode:  401,
				Description: "Invalid token",
				Data:        err.Error(),
			})
		}
		role := claims["role"].(string)

		allow, err := casbinPermission.checkPermission(ctx, role)
		if err != nil {
			return ctx.Status(401).JSON(models.Response{
				StatusCode:  401,
				Description: "error while enforcing",
				Data:        err.Error(),
			})
		}
		if !allow {
			return ctx.Status(401).JSON(models.Response{
				StatusCode:  401,
				Description: "You don't have right permission",
				Data:        nil,
			})
		}

		return ctx.Next()
	}

}

func (c *casbinPermission) checkPermission(ctx *fiber.Ctx, role string) (bool, error) {
	subject := role
	object := ctx.OriginalURL()
	action := ctx.Method()

	allow, err := c.enforcer.Enforce(subject, object, action)
	if err != nil {
		return false, err
	}

	return allow, nil
}
