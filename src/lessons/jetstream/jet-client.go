package main

import (
	"context"
	"flag"
	"fmt"
	// "math/rand"

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
	ctx := context.Background()

	stream, _ := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{Name: "LOGOS", Subjects: []string{"LOGOS.>"}})
	fmt.Println("\nLOGOS, created stream: ", stream)

	app.Get("/bob/", func(c *fiber.Ctx) error {
		_, pubErr := js.Publish(ctx, "LOGOS.msg", []byte("LOGOS, publishing a message with counter: "+fmt.Sprint(counter)))
		if pubErr != nil {
			fmt.Println("\nLOGOS, error publishing message: ", pubErr)
		} else {
			fmt.Println("\nLOGOS, publishing a message with counter: ", counter)
			counter++
		}
		return c.SendString("bob" + fmt.Sprint(counter))
	})

	// rand_num := rand.Int() // create random number
	consumer, consErr := js.CreateOrUpdateConsumer(ctx, "LOGOS", jetstream.ConsumerConfig{})

	if consErr != nil {
		fmt.Println("\nLOGOS, error creating consumer: ", consErr)
	} else {
		fmt.Println("\nLOGOS, created consumer: ", consumer)
		_, consumeErr := consumer.Consume(func(msg jetstream.Msg) {
			fmt.Println("\nLOGOS, received message: ", string(msg.Data()))
		})

		if consumeErr != nil {
			fmt.Println("\nLOGOS, error consuming messages: ", consumeErr)
		}
	}

	var port int
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.Parse()
	app.Listen(":" + fmt.Sprint(port))
}
