package main

import (
	mux "github.com/h3isenbug/url-common/pkg/event-mux"
	consumer "github.com/h3isenbug/url-common/pkg/message-queue-consumer"
	"github.com/h3isenbug/url-query/config"
	url2 "github.com/h3isenbug/url-query/recipients/url"
	"github.com/streadway/amqp"
)

func provideAMQPChannel() (*amqp.Channel, func(), error) {
	con, err := amqp.Dial(config.Config.RabbitServer)
	if err != nil {
		return nil, nil, err
	}

	channel, err := con.Channel()

	return channel, func() {
		channel.Close()
		con.Close()
	}, err
}

func provideMessageConsumer(channel *amqp.Channel, messageMux mux.MessageMux) (consumer.MessageQueueConsumer, func(), error) {
	var cons, err = consumer.NewRabbitMQQueueConsumerV1(channel, func(tag uint64) { channel.Ack(tag, false) }, messageMux, config.Config.RabbitQueueName)
	return cons, func() { cons.GracefulShutdown() }, err
}

func provideReadRecipient(channel *amqp.Channel) (url2.ReadRecipient, error) {
	return url2.NewRabbitMQReadRecipient(channel, config.Config.RabbitQueueName)
}
