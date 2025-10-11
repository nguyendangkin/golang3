package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func BodyParserAndValidate(c *fiber.Ctx, v any) error {
	if err := c.BodyParser(v); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
		return err
	}

	if err := validate.Struct(v); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			errorList := []fiber.Map{}
			for _, e := range errs {
				errorList = append(errorList, fiber.Map{
					"field":   e.Field(),
					"tag":     e.Tag(),
					"message": fmt.Sprintf("%s failed on %s", e.Field(), e.Tag()),
				})
			}
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errorList,
			})
			return err
		}
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	return nil
}
