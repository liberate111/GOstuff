package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Respond with "Hello, World!" on root path, "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// GET http://localhost:8080/hello%20world
	app.Get("/value/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// GET http://localhost:3000/john
	app.Get("/name/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	// GET http://localhost:3000/api/user/john

	app.Get("/path/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	// GET http://localhost:3000/static/hello.html
	app.Static("/static", "./public")

	app.Get("/error/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	// Use: Middleware
	// Match any request
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Match request starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Attach multiple handlers
	app.Use("/api2", func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "random.String(32)")
		return c.Next()
	}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Mount
	micro := fiber.New()
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("john doe")
	})
	app.Mount("/john", micro) // GET /john/doe -> 200 OK

	// Group
	// api := app.Group("/api", handler) // /api

	// v1 := api.Group("/v1", handler) // /api/v1
	// v1.Get("/list", handler)        // /api/v1/list
	// v1.Get("/user", handler)        // /api/v1/user

	// v2 := api.Group("/v2", handler) // /api/v2
	// v2.Get("/list", handler)        // /api/v2/list
	// v2.Get("/user", handler)        // /api/v2/user

	app.Listen(":3000")
}
