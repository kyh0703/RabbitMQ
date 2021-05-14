package rabbit

import (
	"sync"

	"github.com/streadway/amqp"
)

var instance *amqp.Connection
var once sync.Once

func Init(url string) error {
	con, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	instance = con
	return nil
}

func GetConnection() *amqp.Connection {
	return instance
}
