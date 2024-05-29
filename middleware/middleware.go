package middleware

import "github.com/gofiber/fiber/v2"

//jwt미들웨어

func JWTMiddleware(c *fiber.Ctx) error {
	// tokenString := c.Get("X-Authorization")

	c.Locals("user_id", 8)

	return c.Next()
}
