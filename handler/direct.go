package handler

import (
	"log"

	rb "../rabbit"
	"github.com/streadway/amqp"
)


type onRecv func(int) string
type Direct struct {
	Queue    string
	RouteKey string
}

func NewDirect(queue string, routeKey string) *Direct {
	return &Direct{
		Queue: queue,
		RouteKey: routeKey,
	}
}

func Close() {

}

func (r *Direct) Send() (error) {
	var conn *amqp.Connection
	conn = rb.GetConnection()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	body := "Hello World"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return true, nil
}

func (r *Direct) Receive(fn onRecv) (bool, error) {
	conn, err := amqp.Dial("amqp://kyh0703:1234@localhost:5672/")
	if err != nil {
		return false, err;
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return false, err;
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
	return true, err
}