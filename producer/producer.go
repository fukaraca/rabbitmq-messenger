package producer

import (
	"github.com/streadway/amqp"
	"log"
)

func Connect(room string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("failed to connect rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to connect channel")
	}

	err = ch.ExchangeDeclare("chat", "direct", false, true, false, true, nil)
	if err != nil {
		log.Fatalf("exchange declaration failed")
	}

	_, err = ch.QueueDeclare(room, false, true, false, true, nil)
	if err != nil {
		log.Fatalf("queue declaration failed")
	}

	err = ch.QueueBind(room, "message", "chat", true, nil)
	if err != nil {
		log.Fatalf("queue declaration failed")
	}
	return ch
}

func Send(ch *amqp.Channel, room, nick string) {

}
