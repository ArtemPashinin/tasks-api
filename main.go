package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It's work!")
	})

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal("При запуске сервреа произошла ошибка")
	}
}