package users

import (
	"natasha/src/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := utils.BodyParserAndValidate(c, &req); err != nil {
		return nil
	}

	user := User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := handleRegisterUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"error": "Failed to register",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(user)

}
