package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// NewRabbitMQConnection sets up a RabbitMQ connection and returns the connection and channel.
func NewRabbitMQConnection() (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

// CloseConnection closes the RabbitMQ connection and channel.
func CloseConnection(conn *amqp.Connection, ch *amqp.Channel) {
	if ch != nil {
		if err := ch.Close(); err != nil {
			log.Printf("Error closing channel: %s", err)
		}
	}
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %s", err)
		}
	}
}
