package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("LOGOS!!!!")
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		c.WriteJSON(fiber.Map{"msg": "LOGOS, i AGAPE YOU!!!! (from server)"})
	}))

	app.Listen(":8080")
}
