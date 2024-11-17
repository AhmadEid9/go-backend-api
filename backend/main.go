package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Ninja struct {
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
}

var ninja Ninja

func getNinja(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	ninja = Ninja{Name: body.Name, Weapon: body.Weapon}
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	ninjaApp := app.Group("/ninja")

	ninjaApp.Get("/", getNinja)

	ninjaApp.Post("/", createNinja)
	log.Fatal(app.Listen(":3000"))
}
