package routes

import (
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/configs"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func SetupRoutes(app *fiber.App) {

	// api := app.Group("/api", jwtware.New(configs.ConfigAuth)) //  api
	auth := app.Group("/api/v1/auth")
	products := app.Group("/api/v1/products", jwtware.New(configs.ConfigAuth)) // auth

	// Product route
	products.Get("/", controllers.GetProductsAll)
	products.Get("/:id", controllers.GetProduct)
	products.Post("/", controllers.CreateProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	// auth route
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

}
