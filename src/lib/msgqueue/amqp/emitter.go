package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connecttion *amqp.Connection
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connecttion.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
}

func NewAMQPEventEmitter(conn *amqp.Connection) (*amqpEventEmitter, error) {
	emitter := &amqpEventEmitter{
		connecttion: conn,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return emitter, nil
}

func (a *amqpEventEmitter) Emit(event Event) error {
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	chann, err := a.connecttion.Channel()
	if err != nil {
		return err
	}

	defer chann.Close()

	msg := amqp.Publishing{
		Headers:     amqpTable{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "applicatioin/json",
	}

	return chann.Publish(
		"events",
		event.EventName(),
		false,
		false,
		msg,
	)
}
