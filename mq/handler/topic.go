// package handler

// import (
// 	"log"

// 	"github.com/streadway/amqp"
// )

// type Topic struct {
// 	Url      string
// 	Queue    string
// 	Route string
// }

// func NewTopic(url string, queue string, route string) *Topic {
// 	return &Topic{
// 		Url: url,
// 		Queue: queue,
// 		Route: route,
// 	}
// }

// func (r *Topic) Send() (bool, error) {
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

// func (r *Topic) Receive(fn onRecv) (bool, error) {
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