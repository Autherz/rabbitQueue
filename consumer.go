package myAmqp

import (
	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume(delivery amqp.Delivery)
}