package main

import (
	"cv-website/builder"
	"cv-website/handler"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	app.Use("/assets", static.New("./assets"))

	app.Use(helmet.New())
	app.Use(compress.New())

	builder.RunStaticGenerator()

	h := handler.HomePageHandler{}

	app.Get("/", h.Show)

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{}))
}
