package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	ampq "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrder() Order {
	return Order{
		ID: uuid.New().String(),
		// random value between 0 and 100
		Price: rand.Float64() * 100,
	}
}

func Notify(ch *ampq.Channel, order Order) error {

	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		ampq.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err
}

func main() {
	conn, err := ampq.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	for i := 0; i < 100; i++ {
		order := GenerateOrder()
		err = Notify(ch, order)

		if err != nil {
			panic(err)
		}

		fmt.Println("Order sent: ", order)
	}
}
