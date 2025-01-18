package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafitanujaya/go-fiber-template/src/config"
	"github.com/rafitanujaya/go-fiber-template/src/di"
	userController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/user"
	"github.com/rafitanujaya/go-fiber-template/src/http/routes"
	userroutes "github.com/rafitanujaya/go-fiber-template/src/http/routes/user"
	"github.com/samber/do/v2"
)

type HttpServer struct{}

func (s *HttpServer) Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
	})

	//? Depedency Injection
	//? UserController
	uc := do.MustInvoke[userController.UserControllerInterface](di.Injector)

	routes := routes.SetRoutes(app)
	userroutes.SetRouteUsers(routes, uc)
	app.Listen(config.GetPort())
}
