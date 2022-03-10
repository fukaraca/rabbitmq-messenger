package consumer

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

var q amqp.Queue
var label = color.New(color.BgGreen, color.Bold, color.FgHiWhite)

//Connect connects and returns a new amqp channel with given 'room' named exchange and random named queue
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

	q, err = ch.QueueDeclare("", false, false, false, true, nil)
	if err != nil {
		log.Fatalf("queue declaration failed")
	}

	err = ch.QueueBind(q.Name, "", room, true, nil)
	if err != nil {
		log.Fatalf("queue binding failed")
	}
	return ch
}

//Listen consumes all messages that were subscribed
func Listen(ch *amqp.Channel, nick string) error {

	forever := make(chan bool)

	del, err := ch.Consume(q.Name, "", true, false, false, true, nil)
	if err != nil {
		log.Println("delivery failed:", err)
		return err
	}
	go func() {
		for d := range del {
			s := string(d.Body)
			if !strings.HasPrefix(s, nick+" :") {
				temp := strings.SplitAfterN(s, " :", 2)
				label.Printf("%s", temp[0])
				fmt.Printf("%s\n", temp[1])
			}

		}
	}()
	<-forever
	return nil
}
