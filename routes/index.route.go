package routes

import (
	"github.com/FianGumilar/gofiber-rest-api/config"
	"github.com/FianGumilar/gofiber-rest-api/controllers"
	"github.com/FianGumilar/gofiber-rest-api/middleware"
	"github.com/FianGumilar/gofiber-rest-api/utils"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// Static Files
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	// Login
	r.Post("/login", controllers.Login)

	r.Get("/users", middleware.UserAuth, controllers.GetAllUsers)
	r.Post("/users", controllers.CreateUsers)
	r.Get("/users/:id", controllers.GetUserById)
	r.Patch("/users/:id", controllers.UpdateUser)
	r.Delete("/users/:id", controllers.DeleteUser)

	// Book
	r.Get("/books", controllers.GetAllBook)
	r.Post("/books", utils.SingleFile, controllers.CreateBook)
}
