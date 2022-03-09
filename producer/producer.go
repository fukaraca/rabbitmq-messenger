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

	err = ch.ExchangeDeclare(room, "fanout", false, false, false, true, nil)
	if err != nil {
		log.Fatalf("exchange declaration failed")
	}

	q, err := ch.QueueDeclare("", false, false, false, true, nil)
	if err != nil {
		log.Fatalf("queue declaration failed")
	}

	err = ch.QueueBind(q.Name, "", room, true, nil)
	if err != nil {
		log.Fatalf("queue binding failed")
	}
	return ch
}

func Send(ch *amqp.Channel, room, nick, msg string) error {
	msg = nick + " : " + msg
	err := ch.Publish(room, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	return err
}
