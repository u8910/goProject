package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	app.Use("/test", func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("你hao")
	})

	log.Fatal(app.Listen(":3000"))
}
