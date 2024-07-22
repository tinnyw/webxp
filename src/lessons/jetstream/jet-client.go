package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	app := fiber.New()
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)
	defer nc.Close()

	fmt.Println("LOGOS, i AGAPE YOU!!!! (from client)")

	counter := 0

	// retrieve consumer handle from a stream
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, _ := js.Stream(ctx, "msg.LOGOS")
	cons, _ := stream.Consumer(ctx, "msg.LOGOS")

	app.Get("/bob/", func(c *fiber.Ctx) error {
		js.Publish(ctx, "msg.LOGOS", []byte("LOGOS, publishing a message with counter: "+fmt.Sprint(counter)))
		counter++
		return c.SendString("bob 8")
	})

	cc, _ := cons.Consume(func(msg jetstream.Msg) {
		fmt.Println("\nLOGOS, we got a message!", string(msg.Data()))
		msg.Ack()
	})
	defer cc.Stop()

	var port int
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.Parse()
	app.Listen(":" + fmt.Sprint(port))
}
