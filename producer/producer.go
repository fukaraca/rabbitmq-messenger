package producer

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("failed to connect to rabbitmq")
	}
	defer conn.Close()

}

func init() {

}
