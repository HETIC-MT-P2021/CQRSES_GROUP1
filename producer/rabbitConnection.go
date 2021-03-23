package producer

import (
	"fmt"
	"log"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/consts"
	"github.com/streadway/amqp"
)

// RabbitMQ connection global instance
var RabbitMQ *amqp.Connection
var CommandChannel *amqp.Channel

func queuefailOnError(err error, msg string, queueName string) {
	if err != nil {
		log.Fatalf("%s: %s for queue: %s", msg, err, queueName)
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s ", msg, err)
	}
}

func initQueue(queueName string) {
	_, err := CommandChannel.QueueDeclare(
		queueName, // name
		false,     // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	queuefailOnError(err, "Failed to declare a queue", queueName)
}

// ConnectToRabbit connects to RabbitMQ instance
func ConnectToRabbit(host string, port string, user string, password string) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	instanceTmp, err := amqp.Dial(connectionString)

	numberOfTest := 0

	for err != nil && numberOfTest < 5 {
		fmt.Println(err)
		fmt.Println("Connection to the rabbitMQ did not succeed, new try")

		time.Sleep(5 * time.Second)
		instanceTmp, err = amqp.Dial(connectionString)

		numberOfTest++
	}

	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Connected to RabbitMQ server successfully!")

	channel, err := instanceTmp.Channel()

	failOnError(err, "Failed to open a channel")

	CommandChannel = channel

	allQueues := []string{consts.CREATE_POST_COMMAND_QUEUE, consts.UPDATE_POST_COMMAND_QUEUE, consts.DELETE_POST_COMMAND_QUEUE}

	for _, queueName := range allQueues {
		initQueue(queueName)
	}

	RabbitMQ = instanceTmp
}
