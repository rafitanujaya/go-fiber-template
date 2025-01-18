package userController

import "github.com/gofiber/fiber/v2"

type UserControllerInterface interface {
	Register(C *fiber.Ctx) error
}
