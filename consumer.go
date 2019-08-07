package rabbitQueue

import (
	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume(delivery amqp.Delivery)
	Transcode(inputPath string, outputPath string, resolution string, socketIoFlag bool) error
}