package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"time"
)

func Root(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func main() {
	engine := html.New("./public", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static 설정 / Directory: ./public | Web: /static
	app.Static("/static", "./public", fiber.Static{
		Compress: true,
		ByteRange: true,
		Browse: true,
		Index: "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge: 3600,
	})

	app.Get("/", Root)


	app.Listen(":3000")
}