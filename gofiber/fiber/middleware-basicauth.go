package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {

	app := fiber.New()
	app.Listen(":3000")

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("html")             // "html"
		c.Accepts("text/html")        // "text/html"
		c.Accepts("application/json") // "application/json"
		c.AcceptsCharsets("utf-8", "utf-16", "iso-8859-1")

		c.Append("Link", "http://google.com", "http://localhost")
		// => Link: http://localhost, http://google.com

		return c.SendString(c.BaseURL())
	})
}
