package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()
	app.Use(favicon.New())
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./public"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(c.BaseURL())
	})

	app.Listen(":3000")
}
