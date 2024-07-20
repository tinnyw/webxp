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

	counter := 0

	app.Get("/bob/", func(c *fiber.Ctx) error {
		nc.Publish("msg.LOGOS", []byte("LOGOS, publishing a message with counter: "+fmt.Sprint(counter)))
		counter++
		return c.SendString("bob 8")
	})

	nc.Subscribe("msg.LOGOS", func(m *nats.Msg) {
		fmt.Printf("\nLOGOS, we got a message: " + string(m.Data))
	})

	var port int
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.Parse()
	app.Listen(":" + fmt.Sprint(port))
}
