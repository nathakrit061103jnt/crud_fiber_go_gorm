package routes

import (
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetProductsAll)

	app.Get("/api/products", controllers.GetProductsAll)
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Post("/api/product", controllers.CreateProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)

}
