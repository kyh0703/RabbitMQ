// package handler

// import (
// 	"log"

// 	"github.com/streadway/amqp"
// )

// type PubSub struct {
// 	Queue    string
// 	Route string
// }

// func NewPubSub(queue string, route string) *PubSub {
// 	return &PubSub{
// 		Queue: queue,
// 		Route: route,
// 	}
// }

// func (r *PubSub) Send() (bool, error) {
// 	conn, err := amqp.Dial(r.Url)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return false, err;
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"hello",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	body := "Hello World"
// 	err = ch.Publish(
// 		"",
// 		q.Name,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})

// 	return true, nil
// }

// func (r *PubSub) Receive(fn onRecv) (bool, error) {
// 	conn, err := amqp.Dial("amqp://kyh0703:1234@localhost:5672/")
// 	if err != nil {
// 		return false, err;
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return false, err;
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"hello",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	msgs, err := ch.Consume(
// 		q.Name,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	forever := make(chan bool)
// 	go func() {
// 		for d := range msgs {
// 			log.Printf("Received a message: %s", d.Body)
// 		}
// 	}()

// 	<-forever
// 	return true, err
// }