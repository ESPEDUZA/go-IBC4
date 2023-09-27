package routes

import (
	"github.com/ESPEDUZA/go-IBC4/tree/main/fiberTest/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")
	api.Get("/users", userHandler.GetAllUsers)
	api.Get("/users/:id", userHandler.GetUser)
	api.Post("/users", userHandler.CreateUser)
	api.Patch("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)
}
