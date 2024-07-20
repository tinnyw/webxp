package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	app := fiber.New()
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	fmt.Println("LOGOS, i AGAPE YOU!!!! (from client)")

	app.Get("/bob/", func(c *fiber.Ctx) error {
		nc.Publish("msg.LOGOS", []byte("bob 1"))
		fmt.Println("\nPublishing event")
		return c.SendString("bob 8")
	})

	nc.Subscribe("msg.LOGOS", func(m *nats.Msg) {
		fmt.Printf("\nLOGOS, we got a message!")
	})

	var port int
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.Parse()
	app.Listen(":" + fmt.Sprint(port))
}
