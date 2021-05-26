package routes

import (
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/configs"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", jwtware.New(configs.ConfigAuth)) //  api
	auth := app.Group("/auth")                                // auth

	// Product route
	api.Get("/products", controllers.GetProductsAll)
	api.Get("/product/:id", controllers.GetProduct)
	api.Post("/product", controllers.CreateProduct)
	api.Put("/product/:id", controllers.UpdateProduct)
	api.Delete("/product/:id", controllers.DeleteProduct)

	// auth route
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)

}
