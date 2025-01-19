package userController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rafitanujaya/go-fiber-template/src/exceptions"
	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/request"
	userService "github.com/rafitanujaya/go-fiber-template/src/services/user"
	"github.com/samber/do/v2"
)

type UserController struct {
	userService userService.UserServiceInterface
}

func NewUserController(userService userService.UserServiceInterface) UserControllerInterface {
	return &UserController{userService: userService}
}

func NewUserControllerInject(i do.Injector) (UserControllerInterface, error) {
	_userService := do.MustInvoke[userService.UserServiceInterface](i)
	return NewUserController(_userService), nil
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	userRequestParse := request.UserRegister{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	response, err := uc.userService.Register(context.Background(), userRequestParse)

	if err != nil {
		return err
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(201).JSON(response)
}
