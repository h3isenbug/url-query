package url

import (
	"encoding/json"
	"github.com/h3isenbug/url-common/pkg/messages"
	"github.com/streadway/amqp"
)

type ReadRecipient interface {
	PublishURLViewed(shortPath, etag, userAgent string) error
}

type RabbitMQReadRecipient struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQReadRecipient(channel *amqp.Channel, queueName string) (ReadRecipient, error) {
	var _, err = channel.QueueDeclare(
		queueName,
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQReadRecipient{
		channel:   channel,
		queueName: queueName,
	}, nil
}

func (recipient RabbitMQReadRecipient) PublishURLViewed(shortPath, etag, userAgent string) error {
	var event = messages.URLViewedEvent{
		Type:      "view",
		ShortPath: shortPath,
		ETag:      etag,
		UserAgent: userAgent,
	}

	var buffer, _ = json.Marshal(&event)

	return recipient.channel.Publish(
		"",
		recipient.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        buffer,
		},
	)
}
