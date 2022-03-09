package consumer

import (
	"github.com/fatih/color"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

var subText = color.New(color.BgWhite, color.Bold, color.FgBlack)
var pubText = color.New(color.BgGreen, color.Bold, color.FgBlack)

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

func Listen(ch *amqp.Channel, room, nick string) {

	forever := make(chan bool)

	del, err := ch.Consume(room, "", true, false, false, true, nil)
	if err != nil {
		log.Println("delivery failed:", err)
	}
	go func() {
		for d := range del {
			s := string(d.Body)
			if strings.HasPrefix(s, nick+" :") {
				pubText.Printf("%s\n", s)
			} else {
				subText.Printf("%s\n", d.Body)
			}

		}
	}()
	<-forever
}
