package myAmqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitQueue struct {
	Deliveries <-chan amqp.Delivery
}

func (queue *RabbitQueue) AddConsumer(tag string, consumer Consumer) string {
	fmt.Println(consumer)
	go queue.consumerConsume(consumer)
	return ""
}

func (queue *RabbitQueue) consumerConsume(consumer Consumer) {

	go func() {
		// channel <-
		for delivery := range queue.Deliveries {
			fmt.Println("test")
			consumer.Consume(delivery)
		}
	}()
	// for delivery := range queue.Deliveries {
	// 	consumer.Consume(delivery)
	// 	// delivery.Ack(false)

	// }

}
