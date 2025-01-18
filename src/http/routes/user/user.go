package userroutes

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/user"
)

func SetRouteUsers(router fiber.Router, uc userController.UserControllerInterface) {
	router.Post("/register", uc.Register)
}
