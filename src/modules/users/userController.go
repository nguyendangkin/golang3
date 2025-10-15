package users

import (
	"natasha/src/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var req RegisterUserRequest
	if err := utils.BodyParserAndValidate(c, &req); err != nil {
		return nil
	}

	err := handleRegisterUser(&req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(&req)

}
