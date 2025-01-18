package routes

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) fiber.Router {
	route := app.Group("/v1")

	return route
}
