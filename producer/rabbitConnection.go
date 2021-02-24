package producer

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// RabbitMQ connection global instance
var RabbitMQ *amqp.Connection
var CreatePostQueue *amqp.Queue
var CommandChannel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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

	q, err := channel.QueueDeclare(
		"createPost", // name
		false,   // durable
		true,    // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	RabbitMQ = instanceTmp
	CreatePostQueue = &q
	CommandChannel = channel
}