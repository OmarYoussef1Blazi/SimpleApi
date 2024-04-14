package routers

import (
	"fmt"
	"testApi/controllers"

	"github.com/gofiber/fiber"
)


func SetupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", func(c *fiber.Ctx) {
		c.SendString("Welcome")
	})
	// User endpoints
	app.Post("/order/add", func(c *fiber.Ctx) {
		fmt.Printf("id is %v ",c.Body())
		 controllers.CreateOrder(c);
	})
	 app.Get("/orders", func(c *fiber.Ctx) {
		controllers.GetOrders(c)
	 })
	 app.Get("/orders/:id", func(c *fiber.Ctx) {
		// id := c.Params("id")
		controllers.GetOrder(c)
	 })
	 app.Put("/update/:id",func(c *fiber.Ctx) {
		//id := c.Params("id");
		controllers.UpdateOrder(c)
	 })
	 app.Delete("/delete/:id", func(c *fiber.Ctx) {
		// := c.Params("id")
		controllers.DeleteOrders(c);
	 })
}


