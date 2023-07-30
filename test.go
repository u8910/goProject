package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	app.Post("/login", func(ctx *fiber.Ctx) error {
		value := ctx.FormValue("user")
		psw := ctx.FormValue("psw")

		return ctx.SendString("你好:" + value + " " + psw)
	})

	app.Static("/", "./")

	log.Fatal(app.Listen(":3000"))
}
