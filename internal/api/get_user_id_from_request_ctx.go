package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func GetUserIDFromRequestCtx(c *fiber.Ctx) (string, error) {
	id := c.Locals("user_id")
	userID, ok := id.(string)
	if !ok {
		return "", errors.New("user_id is not a string")
	}

	return userID, nil
}
