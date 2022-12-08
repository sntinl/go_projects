package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Person struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("error connect %v \n", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("error channel %v \n", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("error channel %v \n", err)
	}
	p := Person{
		Name:     "oscar",
		Age:      23,
		Email:    "campos.herrera.oscar",
		Password: "123",
	}
	DataJson, _ := json.Marshal(p)
	err = ch.Publish("",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        DataJson,
		})
	if err != nil {
		fmt.Printf("error channel %v \n", err)
	}

}
