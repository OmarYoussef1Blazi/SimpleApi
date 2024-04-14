package main

import (
	"fmt"
	"log"
	"testApi/routers"

	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("MongoDB API")
	app := fiber.New()
	routers.SetupRoutes(app)
	fmt.Println("Server is getting started...")
	fmt.Println("Listening on PORT 4000")
	log.Fatal(app.Listen(":4000"))
}

