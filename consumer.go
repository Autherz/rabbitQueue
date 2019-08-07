<<<<<<< HEAD
package myAmqp
=======
package rabbitQueue
>>>>>>> dev

import (
	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume(delivery amqp.Delivery)
<<<<<<< HEAD
=======
	Transcode(inputPath string, outputPath string, resolution string, socketIoFlag bool) error
>>>>>>> dev
}