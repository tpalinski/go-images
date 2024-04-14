package rabbit

import (
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var connection *amqp.Connection;

type RabbitHandler interface {
	OnMessage(amqp.Delivery) error
}

func InitRabbitConnection(retries int, timeout int, handler *RabbitHandler) error {
	rabit_addr, ok := os.LookupEnv("RABBIT_ADDR")
	if !ok {
		rabit_addr = "localhost"
	}
	connection_string := fmt.Sprintf("amqp://guest:guest@%s:5672/", rabit_addr)
	log.Info("Connecting to rabbitmq instance... ")
	for i:=0; i<retries; i++ {
		conn , err := amqp.Dial(connection_string);
		if err == nil {
			connection = conn;
			defer connection.Close();
                        log.Infof("Connected to rabbitmq instance at %s", rabit_addr);
			c, err := conn.Channel();
			defer c.Close();
			if err != nil {
				log.Error("Error while creating amqp channel")
				return err
			}
			_, err = c.QueueDeclare("raw", false, false, false, false, nil);
			if err != nil {
				return err
			}
			err = c.ExchangeDeclare("images", "direct", false, false, false, false, nil);
			if err != nil {
				return err
			}
			err = c.QueueBind("raw", "raw_images", "images", false, nil);
			if err != nil {
				return err
			}
			log.Infof("Declared queues and exchanges")
			handleMessages(c, handler)
		} else {
			log.Warnf("Could not connect to rabbitmq, no. of retries: %d", i);
			time.Sleep(time.Duration(timeout * int(time.Second)))
		}
	}
	return amqp.Error{}
}

func handleMessages(channel *amqp.Channel, handler *RabbitHandler) {
	var forever chan struct {};
	log.Info("Starting to listen for messages")
	msgs, _ := channel.Consume(
		"raw",
		"transformer",
		true,
		false,
		false, 
		false,
		nil,
	)
	go func() {
                for range msgs {
			log.Info("Received a message")
		}
	}()

	<-forever;
}

