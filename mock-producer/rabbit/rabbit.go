package rabbit

import (
	"context"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var connection *amqp.Connection;
var channel *amqp.Channel;

func InitRabbitConnection(retries int, timeout int) error {
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
                        log.Infof("Connected to rabbitmq instance at %s", rabit_addr);
			c, err := conn.Channel();
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
			channel = c;
			return nil
		} else {
			log.Warnf("Could not connect to rabbitmq, no. of retries: %d", i);
			time.Sleep(time.Duration(timeout * int(time.Second)))
		}
	}
	return amqp.Error{}
}

func CloseConnection() {
	log.Info("Closing connection to rabbitmq")
	connection.Close()
}

func SendMessage(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return channel.PublishWithContext(ctx, "images", "raw_images",false, false, amqp.Publishing{Body: data})
}
