package service

import (
	"github.com/gofiber/fiber/v2"
)

type Ya struct {}

func NewYa() *Ya {
	y := new(Ya)
	return y
}

func (y *Ya) Run(port string){
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World 👋!")
	})

	app.Listen(port)
}