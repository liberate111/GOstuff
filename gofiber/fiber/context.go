package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func main() {
	app := fiber.New()
	app.Use(favicon.New())

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("html")             // "html"
		c.Accepts("text/html")        // "text/html"
		c.Accepts("application/json") // "application/json"
		c.AcceptsCharsets("utf-8", "utf-16", "iso-8859-1")

		c.Append("Link", "http://google.com", "http://localhost")
		// => Link: http://localhost, http://google.com

		return c.SendString(c.BaseURL())
	})

	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	// curl -X POST http://localhost:8080 -d user=john
	app.Post("/", func(c *fiber.Ctx) error {
		c.FormValue("name")
		// Get raw body from POST request:
		return c.Send(c.Body()) // []byte("user=john")
	})

	// curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"john\",\"pass\":\"doe\"}" localhost:3000
	app.Post("/json", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}
		log.Println(p.Name) // john
		log.Println(p.Pass) // doe
		return c.JSON(p)
	})

	app.Get("/set", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "randomvalue",
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "token",
			// Set expiry date to the past
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./public/hello.html")
		// => Download report-12345.pdf
	})

	app.Get("/req", func(c *fiber.Ctx) error {
		c.Request().Header.Method()
		// => []byte("GET")
		// c.Response().Write([]byte("response!")) // args is type *bufio.Writer

		c.Get("Content-Type") // "text/plain"
		c.Is("text/html")
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/coffee", func(c *fiber.Ctx) error {
		return c.Redirect("/teapot")
	})

	app.Get("/teapot", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("🍵 short and stout 🍵")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		r := c.Route()
		fmt.Println(r.Method, r.Path, r.Params, r.Handlers)
		// GET /hello/:name handler [name]
		s := fmt.Sprintln(r.Method, r.Path, r.Params, r.Handlers)

		c.Set("Content-Type", "text/plain")
		return c.SendString(s)
	})

	app.Listen(":3000")
}
