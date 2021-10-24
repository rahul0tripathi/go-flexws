package amqp

import (
	"fmt"
	"github.com/rahul0tripathi/fastws/internal/handler"
	"github.com/streadway/amqp"
	"log"
)

var (
	amqpConn *amqp.Connection
	channel  *amqp.Channel
	delivery <-chan amqp.Delivery
	_queue   amqp.Queue
	err      error
)

func failOnError(err error) {
	if err != nil {
		log.Fatalf("%s:", err)
	}
}
func RoomQueueHandler(username string, password string, host string, port string, queue string, bindKey []string, exchange string, exchangeType string, consumer string) {
	amqpConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	failOnError(err)
	defer amqpConn.Close()
	channel, err = amqpConn.Channel()
	failOnError(err)
	defer channel.Close()

	err = channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err)
	_queue, err = channel.QueueDeclare(
		queue,
		false,
		true,
		true,
		false,
		nil,
	)
	failOnError(err)
	for _, k := range bindKey {
		err = channel.QueueBind(
			_queue.Name,
			k,
			exchange,
			false,
			nil)
	}
	failOnError(err)
	delivery, err = channel.Consume(
		_queue.Name,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err)
	forever := make(chan bool)
	go func() {
		for d := range delivery {
			handler.MessageHandler(d)
		}
	}()
	<-forever
}
