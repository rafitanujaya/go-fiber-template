package httpServer

import (
	"fmt"

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
	fmt.Printf("New Fiber\n")
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
	})

	fmt.Printf("Inject Controllers\n")
	//? Depedency Injection
	//? UserController
	uc := do.MustInvoke[userController.UserControllerInterface](di.Injector)

	fmt.Printf("Set Routes\n")
	routes := routes.SetRoutes(app)
	fmt.Printf("Set User Route\n")
	userroutes.SetRouteUsers(routes, uc)

	fmt.Printf("Start Lister\n")
	app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort()))
}
