package main

import (
	"fmt"
	service "productService/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, I'm product microservice!")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		data := service.GetList()
		return c.Status(200).JSON(data)
	})

	fmt.Println("🚀 product microservice running...")
	app.Listen(":5002")
}
