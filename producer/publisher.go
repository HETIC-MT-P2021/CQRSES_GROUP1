package producer

import (
	"encoding/json"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/consts"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1/models"
	"github.com/streadway/amqp"
)

// PublishCreatePost sends createPost command to rabbit
func PublishCreatePost(post models.Post) error {
	body, err := json.Marshal(post)

	if err != nil {
		return err
	}

	err = CommandChannel.Publish(
		"",                               // exchange
		consts.CREATE_POST_COMMAND_QUEUE, // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return nil
}

// PublishUpdatePost sends updatePost command to rabbit
func PublishUpdatePost(post models.Post) error {
	body, err := json.Marshal(post)

	if err != nil {
		return err
	}

	err = CommandChannel.Publish(
		"",                               // exchange
		consts.UPDATE_POST_COMMAND_QUEUE, // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return nil
}

// PublishDeletePost sends deletePost command to rabbit
func PublishDeletePost(post models.Post) error {
	body, err := json.Marshal(post)

	if err != nil {
		return err
	}

	err = CommandChannel.Publish(
		"",                               // exchange
		consts.DELETE_POST_COMMAND_QUEUE, // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return nil
}
